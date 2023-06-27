package mainframe

const (
	LOGO_TEXT_ONLY = iota
	LOGO_IMAGE_ONLY
	LOGO_BOTH
)

func GetLogo(mode byte) []string {
	image := []string{
		"\033[48;2;88;149;152m\033[38;2;70;155;155m▄\033[48;2;96;163;181m\033[38;2;75;151;161m▄\033[48;2;127;184;216m\033[38;2;123;193;230m▄\033[48;2;111;172;199m\033[38;2;129;192;228m▄\033[48;2;117;193;230m\033[38;2;120;190;224m▄\033[48;2;124;209;252m\033[38;2;98;178;207m▄\033[48;2;90;157;177m\033[38;2;68;141;151m▄\033[48;2;121;180;209m\033[38;2;116;176;203m▄\033[48;2;119;190;224m\033[38;2;118;188;223m▄\033[48;2;114;195;231m\033[38;2;112;197;235m▄\033[48;2;126;201;242m\033[38;2;126;212;255m▄\033[48;2;95;164;186m\033[38;2;76;153;168m▄\033[48;2;90;157;172m\033[38;2;80;150;161m▄\033[48;2;82;150;160m\033[38;2;91;156;171m▄\033[48;2;74;143;151m\033[38;2;96;158;178m▄\033[48;2;124;183;214m\033[38;2;127;182;213m▄\033[48;2;99;151;161m\033[38;2;89;163;173m▄\033[0m",
		"\033[48;2;78;161;169m\033[38;2;125;210;253m▄\033[48;2;84;164;185m\033[38;2;124;209;252m▄\033[48;2;123;210;254m\033[38;2;127;213;255m▄\033[48;2;128;196;234m\033[38;2;128;195;231m▄\033[48;2;102;166;190m\033[38;2;128;185;218m▄\033[48;2;72;142;151m\033[38;2;115;175;203m▄\033[48;2;75;148;153m\033[38;2;67;143;158m▄\033[48;2;121;180;205m\033[38;2;116;175;212m▄\033[48;2;96;159;177m\033[38;2;97;165;193m▄\033[48;2;77;154;169m\033[38;2;66;148;162m▄\033[48;2;124;208;244m\033[38;2;73;154;171m▄\033[48;2;77;156;168m\033[38;2;66;140;150m▄\033[48;2;107;170;192m\033[38;2;106;168;193m▄\033[48;2;132;186;220m\033[38;2;113;174;201m▄\033[48;2;122;181;211m\033[38;2;71;142;153m▄\033[48;2;122;180;209m\033[38;2;75;149;158m▄\033[48;2;89;162;171m\033[38;2;72;152;157m▄\033[0m",
		"\033[48;2;103;177;203m\033[38;2;90;159;173m▄\033[48;2;123;191;226m\033[38;2;122;179;208m▄\033[48;2;89;167;188m\033[38;2;70;145;152m▄\033[48;2;85;156;171m\033[38;2;66;143;144m▄\033[48;2;103;167;195m\033[38;2;94;163;186m▄\033[48;2;135;185;194m\033[38;2;144;188;190m▄\033[48;2;166;184;53m\033[38;2;204;200;2m▄\033[48;2;177;191;70m\033[38;2;208;205;0m▄\033[48;2;171;188;64m\033[38;2;210;207;0m▄\033[48;2;164;184;50m\033[38;2;210;206;0m▄\033[48;2;165;185;58m\033[38;2;204;200;5m▄\033[48;2;120;174;183m\033[38;2;117;172;169m▄\033[48;2;122;195;237m\033[38;2;120;200;246m▄\033[48;2;109;187;219m\033[38;2;112;197;232m▄\033[48;2;73;151;163m\033[38;2;75;153;165m▄\033[48;2;110;194;230m\033[38;2;127;201;242m▄\033[48;2;115;199;236m\033[38;2;113;193;226m▄\033[0m",
		"\033[48;2;99;167;188m\033[38;2;125;192;228m▄\033[48;2;122;188;221m\033[38;2;125;211;255m▄\033[48;2;87;164;182m\033[38;2;128;212;255m▄\033[48;2;75;155;159m\033[38;2;87;164;182m▄\033[48;2;90;161;181m\033[38;2;64;141;154m▄\033[48;2;132;179;176m\033[38;2;94;155;134m▄\033[48;2;196;196;10m\033[38;2;197;197;8m▄\033[48;2;211;212;2m\033[38;2;208;210;2m▄\033[48;2;215;216;3m\033[38;2;213;214;3m▄\033[48;2;210;210;2m\033[38;2;209;210;3m▄\033[48;2;195;196;11m\033[38;2;195;196;12m▄\033[48;2;83;152;133m\033[38;2;93;164;145m▄\033[48;2;97;182;214m\033[38;2;71;157;167m▄\033[48;2;106;188;218m\033[38;2;72;152;160m▄\033[48;2;75;149;159m\033[38;2;77;149;163m▄\033[48;2;124;180;212m\033[38;2;124;183;212m▄\033[48;2;86;156;169m\033[38;2;91;162;173m▄\033[0m",
		"\033[48;2;112;173;200m\033[38;2;89;162;171m▄\033[48;2;85;164;186m\033[38;2;74;154;162m▄\033[48;2;129;214;255m\033[38;2;95;177;202m▄\033[48;2;125;192;224m\033[38;2;93;162;186m▄\033[48;2;115;177;211m\033[38;2;110;179;215m▄\033[48;2;132;180;176m\033[38;2;130;198;222m▄\033[48;2;208;202;0m\033[38;2;133;171;97m▄\033[48;2;209;203;0m\033[38;2;130;172;90m▄\033[48;2;207;203;0m\033[38;2;147;182;113m▄\033[48;2;205;200;0m\033[38;2;162;195;140m▄\033[48;2;202;197;0m\033[38;2;158;205;157m▄\033[48;2;84;148;131m\033[38;2;110;171;186m▄\033[48;2;94;177;210m\033[38;2;120;199;241m▄\033[48;2;122;206;241m\033[38;2;123;208;249m▄\033[48;2;120;202;241m\033[38;2;96;177;203m▄\033[48;2;131;188;222m\033[38;2;94;161;181m▄\033[48;2;88;158;169m\033[38;2;100;164;183m▄\033[0m",
		"\033[48;2;73;158;155m\033[38;2;75;159;160m▄\033[48;2;76;159;163m\033[38;2;81;164;178m▄\033[48;2;84;165;181m\033[38;2;126;211;255m▄\033[48;2;78;157;178m\033[38;2;128;212;255m▄\033[48;2;104;185;221m\033[38;2;128;212;255m▄\033[48;2;118;205;247m\033[38;2;117;200;236m▄\033[48;2;64;147;170m\033[38;2;77;159;165m▄\033[48;2;63;151;164m\033[38;2;78;161;160m▄\033[48;2;87;162;189m\033[38;2;72;151;156m▄\033[48;2;113;184;232m\033[38;2;86;163;183m▄\033[48;2;119;211;255m\033[38;2;128;210;251m▄\033[48;2;117;183;216m\033[38;2;77;155;174m▄\033[48;2;116;192;225m\033[38;2;71;151;160m▄\033[48;2;99;183;210m\033[38;2;73;155;161m▄\033[48;2;69;148;156m\033[38;2;81;163;174m▄\033[48;2;81;160;179m\033[38;2;126;211;254m▄\033[48;2;116;184;215m\033[38;2;126;210;254m▄\033[0m",
		"\033[48;2;89;176;191m\033[38;2;107;159;174m▄\033[48;2;122;206;246m\033[38;2;124;211;254m▄\033[48;2;129;214;255m\033[38;2;94;174;200m▄\033[48;2;89;170;195m\033[38;2;104;170;195m▄\033[48;2;73;152;167m\033[38;2;87;156;172m▄\033[48;2;87;168;189m\033[38;2;79;159;174m▄\033[48;2;122;206;245m\033[38;2;128;212;255m▄\033[48;2;76;158;172m\033[38;2;109;193;225m▄\033[48;2;96;178;205m\033[38;2;98;180;209m▄\033[48;2;128;213;253m\033[38;2;102;180;209m▄\033[48;2;124;210;255m\033[38;2;126;195;232m▄\033[48;2;73;149;168m\033[38;2;105;170;197m▄\033[48;2;102;183;215m\033[38;2;119;196;235m▄\033[48;2;108;190;224m\033[38;2;121;206;248m▄\033[48;2;71;145;160m\033[38;2;105;182;215m▄\033[48;2;124;187;221m\033[38;2;125;184;217m▄\033[48;2;93;168;187m\033[38;2;103;141;152m▄\033[0m",
	}
	text := []string{
		"\033[1;34;40m┌─────────────────────────────────────────────────┐ \033[0m",
		"\033[1;34;40m|   ██████  ███    ███ ███████  ██████   █████    | \033[0m",
		"\033[1;31;40m|█████████████████████████████████████████████████| \033[0m",
		"\033[0;31;40m|█████████████████████████████████████████████████| \033[0m",
		"\033[1;31;40m|█████████████████████████████████████████████████| \033[0m",
		"\033[1;34;40m|   ██████  ██      ██ ███████  ██████  ██   ██   | \033[0m",
		"\033[1;34;40m└─────────────────────────────────────────────────┘ \033[0m",
		"\033[1;31;40m警告：您目前使用的是Dev预览版本，请勿用于实际生产环境！\033[0m",
		"\033[1;31;40m警告：您目前使用的是Dev预览版本，请勿用于实际生产环境！\033[0m",
		"\033[1;31;40m警告：您目前使用的是Dev预览版本，请勿用于实际生产环境！\033[0m",
	}
	if mode == LOGO_TEXT_ONLY {
		return text
	} else if mode == LOGO_IMAGE_ONLY {
		return image
	} else {
		o := []string{}
		for i := 0; i < 7; i++ {
			o = append(o, image[i]+"\u001B[1;34;40m  \u001B[0m"+text[i])
		}
		return o
	}

}
