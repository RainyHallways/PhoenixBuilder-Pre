package blockNBT_API

import (
	"fmt"
	"phoenixbuilder/fastbuilder/mcstructure"
	"phoenixbuilder/minecraft/protocol/packet"
)

// 使用铁砧修改物品名称时会被使用的结构体
type AnvilChangeItemName struct {
	Slot uint8  // 被修改物品在背包所在的槽位
	Name string // 要修改的目标名称
}

// 在 pos 处放置一个方块状态为 blockStates 的铁砧，并依次发送 request 列表中的物品名称修改请求。
// 当 needToDestroyAnvil 为真时，被放置的铁砧会被移除为空气，否则予以保留。
// 返回值 []bool 代表 request 中每个请求的操作结果，它们一一对应，且为真时代表成功改名。
// 因为如果改名时游戏模式不是创造，或者经验值不足，或者提供的新物品名称与原始值相同，
// 那么都会遭到租赁服的拒绝。但这显然不是一个会导致程序崩溃的错误，所以我们使用布尔值表来描述操作结果。
// 当然，此函数在执行时会自动更换客户端的游戏模式为创造，因此您无需再手动操作一次游戏模式。
// 注：如果这个函数返回了错误，那么您应当以 panic 结束 PhoenixBuilder ，因为只有本函数正常执行时，
// 相应的容器资源才会得到正确释放。
// 当然，欢迎你解决这个问题
// (这个问题还是蛮复杂的，因为不适当的处理可能造成更加深层次的问题，比如客户端的背包数据不正确等)
func (g *GlobalAPI) ChangeItemNameByUsingAnvil(
	pos [3]int32,
	blockStates string,
	request []AnvilChangeItemName,
	needToDestroyAnvil bool,
) ([]bool, error) {
	ans := []bool{}
	// 初始化
	err := g.SendSettingsCommand("gamemode 1", true)
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	// 更换游戏模式为创造
	uniqueId, correctPos, err := g.GenerateNewAnvil(pos, blockStates)
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	// 尝试生成一个铁砧并附带承重方块
	_, err = g.SendWSCommandWithResponce(fmt.Sprintf("tp %d %d %d", correctPos[0], correctPos[1], correctPos[2]))
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	// 传送机器人到铁砧处
	_, lockDown := g.PacketHandleResult.ContainerResources.Occupy(false)
	// 获取容器资源
	got, err := mcstructure.ParseStringNBT(blockStates, true)
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	blockStatesMap, normal := got.(map[string]interface{})
	if !normal {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: Could not convert got into map[string]interface{}; got = %#v", got)
	}
	// 获取要求放置的铁砧的方块状态
	err = g.ChangeSelectedHotbarSlot(0, true)
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	// 切换手持物品栏
	err = g.OpenContainer(correctPos, "minecraft:anvil", blockStatesMap, 0, false)
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	// 打开铁砧
	for _, value := range request {
		datas, err := g.PacketHandleResult.Inventory.GetItemStackInfo(0, value.Slot)
		if err != nil {
			return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
		}
		if datas.Stack.ItemType.NetworkID == 0 {
			continue
		}
		// 获取被改物品的相关信息
		resp, err := g.MoveItem(
			MoveItemDatas{
				WindowID:    0,
				ContainerID: 12,
				Slot:        value.Slot,
			},
			MoveItemDatas{
				WindowID:                  -1,
				ItemStackNetworkIDProvide: 0,
				ContainerID:               0,
				Slot:                      1,
			},
			uint8(datas.Stack.Count),
		)
		if err != nil {
			return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
		}
		if resp[0].Status != 0 {
			return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: Operation %v have been canceled by error code %v; inventorySlot = %v, containerSlot = 1, moveCount = %v", resp[0].RequestID, resp[0].Status, value.Slot, datas.Stack.Count)
		}
		// 移动物品到铁砧
		err = g.WritePacket(&packet.AnvilDamage{
			Damage:        0,
			AnvilPosition: pos,
		})
		if err != nil {
			return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
		}
		// 请求损坏当前铁砧
		successStates, err := g.ChangeItemName(resp[0], value.Name, value.Slot)
		if err != nil {
			return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
		}
		ans = append(ans, successStates)
		// 发送改名请求
	}
	err = g.CloseContainer()
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	// 关闭铁砧
	lockDown.Unlock()
	// 释放容器公用资源
	if needToDestroyAnvil {
		_, err := g.SendWSCommandWithResponce(fmt.Sprintf("setblock %d %d %d air", correctPos[0], correctPos[1], correctPos[2]))
		if err != nil {
			return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
		}
	}
	// 如果需要移除铁砧
	err = g.RevertBlockUnderAnvil(uniqueId, correctPos)
	if err != nil {
		return []bool{}, fmt.Errorf("ChangeItemNameByUsingAnvil: %v", err)
	}
	// 恢复铁砧下方的承重方块为原本方块
	return ans, nil
	// 返回值
}
