package constants

type EventName string

const (
	EventBasketball_free EventName = "basketball-free"
	EventSportStacking   EventName = "sport_stacking"
	EventJuggling_3      EventName = "juggling-3"
	EventJumpRope_2      EventName = "jump_rope-2"
	EventSoccer_lifting  EventName = "soccer-lifting"
	EventVolleyball_toss EventName = "volleyball-toss"
	EventBadminton_wall  EventName = "badminton-wall"
	EventJumpRopeSpeed   EventName = "jump-rope-speed"
)

var EventNameEnToKo = map[EventName]string{
	"sport_stacking":  "스포츠스태킹-사이클",
	"juggling-3":      "저글링-3볼 캐스케이드",
	"jump_rope-2":     "줄넘기-2단 뛰기",
	"jump-rope-speed": "줄넘기-번갈아뛰기",
	"soccer-lifting":  "축구-양발리프팅",
	"badminton-wall":  "배드민턴-벽치기",
}

var Rule = map[string][]string{
	"badminton-wall": {
		"30초 동안 실시한 셔틀콕 벽치기 개수를 합산합니다.",
		"땅에 떨어뜨렸을 경우 주워서 바로 다시 시작해야 해요.(다른 셔틀콕 사용 안됨)",
		"셔틀콕 벽치기 시 높이와 거리는 제한이 없어요.",
		"벽의 모양과 형태 및 장소는 자유롭게 선택할 수 있어요.",
		"벽의 모양과 형태 및 장소는 자유롭게 선택할 수 있어요.",
	},
	"sport_stacking": {
		"공인기록장비로 측정 후 결과값을 촬영해 주세요.",
		"3-6-3 스택, 6-6 스택, 1-10-1 스택이 차례대로 결합하여 반드시 3-6-3이 다운 스태킹 된 상태로 경기를 마무리하여야 해요.",
		"경기 시작 전/후 손의 위치는 기록 장비(타이머)에 위치해야 해요",
		"잘못된 순서로 스태킹을 하거나 양손으로 두 더미를 동시 진행하면 규칙 위반에 해당 되요.",
		"타이머 리셋 상태를 카메라에 비춰준 뒤, 동작을 수행하는 모습 전체가 보이도록 연속으로 촬영해 주세요.",
	},
	"jump_rope-2": {
		"30초간 이중 뛰기 실시한 모든 개수를 측정합니다.",
		"줄을 넘다가 걸리면 바로 이어서 넘어요.",
		"시작 신호 전 줄이 움직이면 안돼요.",
		"동작을 수행하는 모습 전체가 보이도록 연속으로 촬영해 주세요.",
	},
	"jump-rope-speed": {
		"30초 동안 번갈아 뛰기한 개수를 측정합니다.",
		"한 발이 줄을 넘을때마다 한번씩 세요.",
		"줄을 넘다가 걸리면 바로 이어서 넘어요.",
		"시작 전에 줄이 움직이면 안돼요.",
		"동작을 수행하는 모습 전체가 보이도록 연속으로 촬영해 주세요.",
	},
	"soccer-lifting": {
		"30초 동안 팔과 손을 제외한 신체를 사용하여 리프팅의 개수를 합산합니다. (단, 발은 번갈아 가며 사용해야 해요- 한발 연속으로 리프팅 되는 볼은 카운팅 안됨)",
		"볼이 지면에 닿은 상태에서 발로 시작하거나 손으로 토스 후 시작이 가능해요.",
		"바닥에 바운드 후 연결되는 리프팅의 개수는 인정해요.",
		"리프팅 중간에 팔 또는 손을 사용하는 경우 남은 시간과 관계없이 종료됩니다.",
		"동작을 수행하는 모습 전체가 보이도록 연속으로 촬영해 주세요.",
	},
	"juggling-3": {
		"10초 동안 제자리에 서서 3볼 캐스케이드를 실시 한 개수를 합산합니다.",
		"한 손에 2개의 볼, 다른 한 손에 1개의 볼을 쥔 상태에서 시작 신호음에 따라 실시 해요.",
		"스타트 후 10초가 되기 전에 볼을 떨어뜨리거나 멈추면 실격처리 되요.",
		"저글링 볼에 손가락이 손바닥 이외 다른 신체 부위가 닿으면 실격처리 되요.",
		"동작을 수행하는 모습 전체가 연속으로 촬영해 주세요.",
	},
}

var Icon = map[string]string{
	"bronze":   "https://bossports.s3.ap-northeast-2.amazonaws.com/icon/bronze.svg",
	"silver":   "https://bossports.s3.ap-northeast-2.amazonaws.com/icon/silver.svg",
	"gold":     "https://bossports.s3.ap-northeast-2.amazonaws.com/icon/gold.svg",
	"platinum": "https://bossports.s3.ap-northeast-2.amazonaws.com/icon/platinum.svg",
	"diamond":  "https://bossports.s3.ap-northeast-2.amazonaws.com/icon/diamond.svg",
	"master":   "https://bossports.s3.ap-northeast-2.amazonaws.com/icon/master.svg",
}

var Guide = map[string]string{
	"sport_stacking":  "632469757",
	"juggling-3":      "632471769",
	"jump_rope-2":     "632469757",
	"jump-rope-speed": "632469242",
	"soccer-lifting":  "632466005",
	"badminton-wall":  "632467995",
}

var TermsOfService = map[string]string{
	`이용 약관`: `제1장 총칙\n제1조(목적) 이 약관은 ㈜아워스포츠네이션(이하 "회사"라고 한다)이 제공하는 “학교스포츠GG” 서비스의 이용과 관련하여 회사와 이용자와의 권리, 의무 및 책임사항 등을 규정함을 목적으로 합니다.\n제2조(정의) 이 약관에서 사용하는 용어의 정의는 다음과 같습니다.\n  1. \"회사\"라 함은 \"학교스포츠GG\" 사업과 관련된 경제활동을 영위하는 자로서 서비스를 제공하는 자를 말합니다.\n  2. \"이용자\"라 함은 \"회사\"의 이 약관에 따라 \"회사\"가 제공하는 서비스를 이용하는 회원 및 비회원을 말합니다.\n  3. \"회원\"이라 함은 \"회사\"와 이용계약을 체결하고 \"이용자\" 아이디(ID)를 부여받은 \"이용자\"로서 \"회사\"의 정보를 지속적으로 제공받으며 \"회사\"가 제공하는 서비스를 지속적으로 이용할 수 있는 자를 말합니다.\n  4. “비회원”이라 함은“회사”와 이용계약을 체결하기 이전의 자를 말합니다.\n  5. \"아이디(ID)\"라 함은 \"회원\"의 식별과 서비스이용을 위하여 \"회원\"이 정하고 \"회사\"가 승인하는 문자 또는 숫자의 조합을 말합니다.\n  6. \"비밀번호(PASSWORD)\"라 함은 \"회원\"이 부여받은 \"아이디\"와 일치되는 \"회원\"임을 확인하고 비밀보호를 위해 \"회원\" 자신이 정한 문자 또는 숫자의 조합을 말합니다.\n  7. “이용계약”이라 함은 “학교스포츠GG” 이용과 관련하여 “학교스포츠GG”와 “회원” 간에 체결하는 계약을 말합니다.\n  8. “해지”라 함은 “회원”이 “이용계약”을 해약하는 것을 말합니다.\n제3조(신원정보 등의 제공) \"회사\"는 이 약관의 내용, 상호, 대표자 성명, 영업소 소재지 주소(소비자의 불만을 처리할 수 있는 곳의 주소를 포함), 전화번호, 모사전송번호, 전자우편주소, 사업자등록번호, 통신판매업 신고번호 및 개인정보관리책임자 등을 이용자가 쉽게 알 수 있도록 온라인 서비스초기화면에 게시합니다. 다만, 약관은 이용자가 연결화면을 통하여 볼 수 있도록 할 수 있습니다.\n제4조(약관의 게시 등) ① \"회사\"는 이 약관을 \"회원\"이 그 전부를 인쇄할 수 있고 거래과정에서 해당 약관의 내용을 확인할 수 있도록 기술적 조치를 취합니다.\n  ② \"회사\"는 \"회원\"이 \"회사\"와 이 약관의 내용에 관하여 질의 및 응답할 수 있도록 기술적 장치를 설치합니다.\n  ③ \"회사\"는 \"회원\"이 약관에 동의하기에 앞서 약관에 정하여져 있는 내용 중 청약철회 등과 같은 중요한 내용을 이용자가 쉽게 이해할 수 있도록 별도의 연결화면 또는 팝업화면 등을 제공하여 \"회원\"의 확인을 구합니다.\n제5조(약관의 개정 등) ① \"회사\"는 온라인 디지털콘텐츠산업 발전법, 전자상거래 등에서의 소비자보호에 관한 법률, 약관의 규제에 관한 법률 등 관련법을 위배하지 않는 범위에서 이 약관을 개정할 수 있습니다.\n  ② \"회사\"가 약관을 개정할 경우에는 적용일자 및 개정사유를 명시하여 현행약관과 함께 서비스초기화면에 그 적용일자 7일 이전부터 적용일 후 상당한 기간동안 공지하고, 기존회원에게는 개정약관을 전자우편주소로 발송합니다. \n  ③ \"회사\"가 약관을 개정할 경우에는 개정약관 공지 후 개정약관의 적용에 대한 \"회원\"의 동의 여부를 확인합니다. \"회원\"이 개정약관의 적용에 동의하지 않는 경우 \"회사\" 또는 \"회원\"은 콘텐츠 이용계약을 해지할 수 있습니다. 이때, \"회사\"는 계약해지로 인하여 \"회원\"이 입은 손해를 배상합니다.\n제6조(약관의 해석) 이 약관에서 정하지 아니한 사항과 이 약관의 해석에 관하여는 온라인 디지털콘텐츠산업 발전법, 전자상거래 등에서의 소비자보호에 관한 법률, 약관의 규제에 관한 법률, 문화체육관광부장관이 정하는 디지털콘텐츠이용자보호지침, 기타 관계법령 또는 상관례에 따릅니다.\n        제2장 회원가입\n제7조(회원가입) ① 회원가입은 \"이용자\"가 약관의 내용에 대하여 동의를 하고 회원가입신청을 한 후 \"회사\"가 이러한 신청에 대하여 승낙함으로써 체결됩니다.\n  ② 회원가입신청서에는 다음 사항을 기재해야 합니다.\n  1. \"회원\"의 전화번호\n  2. \"비밀번호\"\n  3.“회원”의 이름, 성별, 학교명, 학년, 반, 번호\n  ③ \"회사\"는 상기 \"이용자\"의 신청에 대하여 회원가입을 승낙함을 원칙으로 합니다. 다만, \"회사\"는 다음 각 호에 해당하는 신청에 대하여는 승낙을 하지 않을 수 있습니다.\n  1. 가입신청자가 이 약관에 의하여 이전에 회원자격을 상실한 적이 있는 경우\n  2. 실명이 아니거나 타인의 명의를 이용한 경우\n  3. 허위의 정보를 기재하거나, 회사가 제시하는 내용을 기재하지 않은 경우\n  4. 이용자의 귀책사유로 인하여 승인이 불가능하거나 기타 규정한 제반 사항을 위반하며 신청하는 경우\n  ④ \"회사\"는 서비스 관련 설비의 여유가 없거나, 기술상 또는 업무상 문제가 있는 경우에는 승낙을 유보할 수 있습니다.\n  ⑤ 제3항과 제4항에 따라 회원가입신청의 승낙을 하지 아니하거나 유보한 경우, \"회사\"는 이를 신청자에게 알려야 합니다. \"회사\"의 귀책사유 없이 신청자에게 통지할 수 없는 경우에는 예외로 합니다.\n  ⑥ 회원가입계약의 성립 시기는 \"회사\"의 승낙이 \"이용자\"에게 도달한 시점으로 합니다.\n제8조(“학생”의 회원가입에 관한 특칙) ① 초등학교, 중학교, 고등학교 중 하나에 재학중인 \"이용자\"(이하 “학생”이라고 한다)는 개인정보의 수집 및 이용목적에 대하여 충분히 숙지한 후에 “학교스포츠GG 앱으로 회원가입을 신청하고 본인의 개인정보를 제공하여야 합니다.\n  ② “학생”이 게시한 동영상은 “학생”을 담당하는 선생님, “학생”의 부모 등 법정대리인,“학생”이 허가한 제3자 등이 열람할 수 있습니다.\n  ③ 회원가입을 한 “학생”은 “학생”을 담당하는 선생님에 의하여 관리될 수 있습니다.\n제9조(회원정보의 변경) ① \"회원\"은 개인정보관리화면을 통하여 언제든지 자신의 개인정보를 열람하고 수정할 수 있습니다. \n  ② \"회원\"은 회원가입신청 시 기재한 사항이 변경되었을 경우 온라인으로 수정을 하거나 전자우편 기타 방법으로 \"회사\"에 대하여 그 변경사항을 알려야 합니다. \n  ③ 제2항의 변경사항을 \"회사\"에 알리지 않아 발생한 불이익에 대하여 \"회사\"는 책임지지 않습니다. \n제10조(\"회원\"의 \"아이디\" 및 \"비밀번호\"의 관리에 대한 의무) ① \"회원\"의 \"아이디\"와 \"비밀번호\"에 관한 관리책임은 \"회원\"에게 있으며, 이를 제3자가 이용하도록 하여서는 안 됩니다.\n  ② \"회원\"은 \"아이디\" 및 \"비밀번호\"가 도용되거나 제3자에 의해 사용되고 있음을 인지한 경우에는 이를 즉시 \"회사\"에 통지하고 \"회사\"의 안내에 따라야 합니다.\n  ③ 제2항의 경우에 해당 \"회원\"이 \"회사\"에 그 사실을 통지하지 않거나, 통지한 경우에도 \"회사\"의 안내에 따르지 않아 발생한 불이익에 대하여 \"회사\"는 책임지지 않습니다.\n제11조(\"회원\"에 대한 통지) ① \"회사\"가 \"회원\"에 대한 통지를 하는 경우 “학교스포츠GG”앱 또는 “학교스포츠GG”웹사이트를 통하여 시행합니다.\n  ② \"회사\"는 \"회원\" 전체에 대한 통지의 경우 7일 이상 \"회사\"의 게시판에 게시함으로써 제1항의 통지에 갈음할 수 있습니다. 다만, \"회원\" 본인의 거래와 관련하여 중대한 영향을 미치는 사항에 대하여는 제1항의 통지를  합니다.\n제12조(회원탈퇴 및 자격 상실 등) ① \"회원\"은 \"회사\"에 언제든지 탈퇴를 요청할 수 있으며 \"회사\"는 즉시 회원탈퇴를 처리합니다.\n  ② \"회원\"이 다음 각호의 사유에 해당하는 경우, \"회사\"는 회원자격을 제한 및 정지시킬 수 있습니다.\n  1. 가입신청 시에 허위내용을 등록한 경우\n  2. 다른 사람의 \"회사\"의 서비스이용을 방해하거나 그 정보를 도용하는 등 전자상거래 질서를 위협하는 경우\n  3. \"회사\"를 이용하여 법령 또는 이 약관이 금지하거나 공서양속에 반하는 행위를 하는 경우\n  ③ \"회사\"가 회원자격을 제한·정지시킨 후, 동일한 행위가 2회 이상 반복되거나 30일 이내에 그 사유가 시정되지 아니하는 경우 \"회사\"는 회원자격을 상실시킬 수 있습니다. \n  ④ \"회사\"가 회원자격을 상실시키는 경우에는 회원등록을 말소합니다. 이 경우 \"회원\"에게 이를 통지하고, 회원등록 말소 전에 최소한 30일 이상의 기간을 정하여 소명할 기회를 부여합니다.\n  ⑤ 회원탈퇴와 이용계약 해지는 별개로 취급합니다.\n        제3장 서비스이용계약\n제13조(이용계약의 성립 등) ① \"회원\"은 \"회사\"가 제공하는 다음 또는 이와 유사한 절차에 의하여 이용신청을 합니다. \"회사\"는 계약 체결 전에 각 호의 사항에 관하여 \"회원\"이 정확하게 이해하고 실수 또는 착오 없이 거래할 수 있도록 정보를 제공합니다.\n  1. 학생의 경우 동영상의 게시\n  2. 교사의 경우 자신이 담당하는 학생이 게시한 동영상의 열람 및 선택\n  3. 성명, 전화번호 등의 입력\n  4. 약관내용에 대해 \"회사\"가 취한 조치에 관련한 내용에 대한 확인 \n  5. 이 약관에 동의하고 위 제3호의 사항을 확인하거나 거부하는 표시(예, 마우스 클릭) \n  6. 서비스의 이용신청에 관한 확인 또는 \"회사\"의 확인에 대한 동의\n  ② \"회사\"는 \"회원\"의 이용신청이 다음 각 호에 해당하는 경우에는 승낙하지 않거나 승낙을 유보할 수 있습니다.\n  1. 실명이 아니거나 타인의 명의를 이용한 경우\n  2. 허위의 정보를 기재하거나, \"회사\"가 제시하는 내용을 기재하지 않은 경우\n  3. 서비스 관련 설비의 여유가 없거나, 기술상 또는 업무상 문제가 있는 경우\n  ③ \"회사\"의 승낙이 제16조 제1항의 수신확인통지형태로 \"회원\"에게 도달한 시점에 계약이 성립한 것으로 봅니다. \n  ④ \"회사\"의 승낙의 의사표시에는 \"회원\"의 이용신청에 대한 확인 및 서비스제공 가능여부, 이용신청의 정정·취소 등에 관한 정보 등을 포함합니다.\n제14조(수신확인통지·이용신청 변경 및 취소) ① \"회사\"는 \"회원\"의 이용신청이 있는 경우 \"회원\"에게 수신확인통지를 합니다. \n  ② 수신확인통지를 받은 \"회원\"은 의사표시의 불일치 등이 있는 경우에는 수신확인통지를 받은 후 즉시 이용신청 변경 및 취소를 요청할 수 있고, \"회사\"는 서비스제공 전에 \"회원\"의 요청이 있는 경우에는 지체 없이 그 요청에 따라 처리하여야 합니다. 다만, 이미 대금을 지불한 경우에는 청약철회 등에 관한 제27조의 규정에 따릅니다. \n제15조(\"회사\"의 의무) ① \"회사\"는 법령과 이 약관이 정하는 권리의 행사와 의무의 이행을 신의에 좇아 성실하게 하여야 합니다.\n  ② \"회사\"는 \"회원\"이 안전하게 서비스를 이용할 수 있도록 개인정보(신용정보 포함)보호를 위해 보안시스템을 갖추어야 하며 개인정보보호정책을 공시하고 준수합니다.\n  ③ \"회사\"는 \"회원\"가 콘텐츠이용을 수시로 확인할 수 있도록 조치합니다.\n  ④ \"회사\"는 콘텐츠이용과 관련하여 \"회원\"으로부터 제기된 의견이나 불만이 정당하다고 인정할 경우에는 이를 지체없이 처리합니다. 이용자가 제기한 의견이나 불만사항에 대해서는 게시판을 활용하여 그 처리과정 및 결과를 전달합니다.\n  ⑤ \"회사\"는 이 약관에서 정한 의무 위반으로 인하여 \"회원\"이 입은 손해를 배상합니다.\n제16조(\"이용자\"의 의무) ① \"회원\"은 다음 행위를 하여서는 안 됩니다. \n  1. 신청 또는 변경 시 허위내용의 기재\n  2. 타인의 정보도용 \n  3. \"회사\"에 게시된 정보의 변경 \n  4. \"회사\"가 금지한 정보(컴퓨터 프로그램 등)의 송신 또는 게시 \n  5. \"회사\"와 기타 제3자의 저작권 등 지적재산권에 대한 침해 \n  6. \"회사\" 및 기타 제3자의 명예를 손상시키거나 업무를 방해하는 행위 \n  7. 외설 또는 폭력적인 말이나 글, 화상, 음향, 기타 공서양속에 반하는 정보를 \"회사\"의 사이트에 공개 또는 게시하는 행위 \n  8. 기타 불법적이거나 부당한 행위\n  ② \"회원\"은 관계법령, 이 약관의 규정, 이용안내 및 \"콘텐츠\"와 관련하여 공지한 주의사항, \"회사\"가 통지하는 사항 등을 준수하여야 하며, 기타 \"회사\"의 업무에 방해되는 행위를 하여서는 안 됩니다.\n제17조(콘텐츠서비스의 제공 및 중단) ① 콘텐츠서비스는 연중무휴, 1일 24시간 제공함을 원칙으로 합니다.\n  ② \"회사\"는 컴퓨터 등 정보통신설비의 보수점검, 교체 및 고장, 통신두절 또는 운영상 상당한 이유가 있는 경우 콘텐츠서비스의 제공을 일시적으로 중단할 수 있습니다. 이 경우 \"회사\"는 제11조[\"회원\"에 대한 통지]에 정한 방법으로 \"회원\"에게 통지합니다. 다만, \"회사\"가 사전에 통지할 수 없는 부득이한 사유가 있는 경우 사후에 통지할 수 있습니다.\n  ③ \"회사\"는 상당한 이유 없이 콘텐츠서비스의 제공이 일시적으로 중단됨으로 인하여 \"회원\"이 입은 손해에 대하여 배상합니다. 다만, \"회사\"가 고의 또는 과실이 없음을 입증하는 경우에는 그러하지 아니합니다.\n  ④ \"회사\"는 콘텐츠서비스의 제공에 필요한 경우 정기점검을 실시할 수 있으며, 정기점검시간은 서비스제공화면에 공지한 바에 따릅니다.\n  ⑤ 사업종목의 전환, 사업의 포기, 업체 간의 통합 등의 이유로 콘텐츠서비스를 제공할 수 없게 되는 경우에는 \"회사\"는 제11조[\"회원\"에 대한 통지]에 정한 방법으로 \"회원\"에게 통지하고 당초 \"회사\"에서 제시한 조건에 따라 \"회원\"에게 보상합니다.\n제18조(콘텐츠서비스의 변경) ① \"회사\"는 상당한 이유가 있는 경우에 운영상, 기술상의 필요에 따라 제공하고 있는 콘텐츠서비스를 변경할 수 있습니다. \n  ② \"회사\"는 콘텐츠서비스의 내용, 이용방법, 이용시간을 변경할 경우에 변경사유, 변경될 콘텐츠서비스의 내용 및 제공일자 등을 그 변경 전 7일 이상 해당 콘텐츠초기화면에 게시합니다.\n  ③ 제2항의 경우에 변경된 내용이 중대하거나 \"회원\"에게 불리한 경우에는 \"회사\"가 해당 콘텐츠서비스를 제공받는 \"회원\"에게 제11조[\"회원\"에 대한 통지]에 정한 방법으로 통지하고 동의를 받습니다. 이때, \"회사\"는 동의를 거절한 \"회원\"에 대하여는 변경전 서비스를 제공합니다. 다만, 그러한 서비스 제공이 불가능할 경우 계약을 해지할 수 있습니다.\n  ④ \"회사\"는 제1항에 의한 서비스의 변경 및 제3항에 의한 계약의 해지로 인하여 \"회원\"이 입은 손해를 배상합니다.\n제19조(정보의 제공) ① \"회사\"는 \"회원\"이 콘텐츠이용 중 필요하다고 인정되는 다양한 정보를 공지사항 등의 방법으로 \"회원\"에게 제공할 수 있습니다. \n제20조(게시물의 삭제) ① \"회사\"는 게시판에 정보통신망이용촉진 및 정보보호 등에 관한 법률을 위반한 청소년유해매체물이 게시되어 있는 경우에는 이를 지체 없이 삭제합니다.\n  ② \"회사\"가 운영하는 게시판 등에 게시된 정보로 인하여 법률상 이익이 침해된 자는 \"회사\"에게 당해 정보의 삭제 또는 반박내용의 게재를 요청할 수 있습니다. 이 경우 \"회사\"는 지체 없이 필요한 조치를 취하고 이를 즉시 신청인에게 통지합니다.\n제21조(저작권 등의 귀속) ① \"회사\"가 작성한 저작물에 대한 저작권 기타 지적재산권은 \"회사\"에 귀속합니다. \n  ② \"회사\"가 제공하는 서비스 중 제휴계약에 의해 제공되는 저작물에 대한 저작권 기타 지적재산권은 해당 제공업체에 귀속합니다.\n  ③ \"회원\"은 \"회사\"가 제공하는 서비스를 이용함으로써 얻은 정보 중 \"회사\" 또는 제공업체에 지적재산권이 귀속된 정보를 \"회사\" 또는 제공업체의 사전승낙 없이 복제, 전송, 출판, 배포, 방송 기타 방법에 의하여 영리목적으로 이용하거나 제3자에게 이용하게 하여서는 안 됩니다. \n  ④ \"회사\"는 약정에 따라 \"회원\"의 저작물을 사용하는 경우 당해 \"회원\"의 허락을 받습니다.. \n제22조(개인정보보호) ① \"회사\"는 제7조 제2항의 신청서기재사항 이외에 \"회원\"의 콘텐츠이용에 필요한 최소한의 정보를 수집할 수 있습니다. 이를 위해 \"회사\"가 문의한 사항에 관해 \"회원\"은 진실한 내용을 성실하게 고지하여야 합니다.\n  ② \"회사\"가 \"회원\"의 개인 식별이 가능한 \"개인정보\"를 수집하는 때에는 당해 \"회원\"의 동의를 받습니다. \n  ③ \"회사\"는 \"회원\"이 이용신청 등에서 제공한 정보와 제1항에 의하여 수집한 정보를 당해 \"회원\"의 동의 없이 목적 외로 이용하거나 제3자에게 제공할 수 없으며, 이를 위반한 경우에 모든 책임은 \"회사\"가 집니다. 다만, 다음의 경우에는 예외로 합니다. \n  1. 통계작성, 학술연구 또는 시장조사를 위하여 필요한 경우로서 특정 개인을 식별할 수 없는 형태로 제공하는 경우  \n  2. 도용방지를 위하여 본인확인에 필요한 경우 \n  3. 약관의 규정 또는 법령에 의하여 필요한 불가피한 사유가 있는 경우 \n  ④ \"회사\"가 제2항과 제3항에 의해 \"회원\"의 동의를 받아야 하는 경우에는 \"개인정보\"관리책임자의 신원(소속, 성명 및 전화번호 기타 연락처), 정보의 수집목적 및 이용목적, 제3자에 대한 정보제공관련사항(제공받는 자, 제공목적 및 제공할 정보의 내용)등에 관하여 정보통신망이용촉진 및 정보보호 등에 관한 법률 제22조 제2항이 규정한 사항을 명시하고 고지하여야 합니다.\n  ⑤ \"회원\"은 언제든지 제3항의 동의를 임의로 철회할 수 있습니다.\n  ⑥ \"회원\"은 언제든지 \"회사\"가 가지고 있는 자신의 \"개인정보\"에 대해 열람 및 오류의 정정을 요구할 수 있으며, \"회사\"는 이에 대해 지체 없이 필요한 조치를 취할 의무를 집니다. \"회원\"이 오류의 정정을 요구한 경우에는 \"회사\"는 그 오류를 정정할 때까지 당해 \"개인정보\"를 이용하지 않습니다.\n  ⑦ \"회사\"는 개인정보보호를 위하여 관리자를 한정하여 그 수를 최소화하며, 신용카드, 은행계좌 등을 포함한 \"회원\"의 \"개인정보\"의 분실, 도난, 유출, 변조 등으로 인한 \"회원\"의 손해에 대하여 책임을 집니다.\n  ⑧ \"회사\" 또는 그로부터 \"개인정보\"를 제공받은 자는 \"회원\"이 동의한 범위 내에서 \"개인정보\"를 사용할 수 있으며, 목적이 달성된 경우에는 당해 \"개인정보\"를 지체 없이 파기합니다. \n  ⑨ \"회사\"는 정보통신망이용촉진 및 정보보호에 관한 법률 등 관계 법령이 정하는 바에 따라 \"회원\"의 \"개인정보\"를 보호하기 위해 노력합니다. \"개인정보\"의 보호 및 사용에 대해서는 관련법령 및 \"회사\"의 개인정보보호정책이 적용됩니다. \n        제4장 콘텐츠이용계약의 청약철회, 계약해제·해지 및 이용제한\n제23조(\"이용자\"의 청약철회와 계약해제·해지) ① \"회사\"와 \"콘텐츠\"의 이용에 관한 계약을 체결한 \"이용자\"는 수신확인의 통지를 받은 날로부터 7일 이내에는 청약의 철회를 할 수 있습니다. 다만, \"회사\"가 다음 각 호중 하나의 조치를 취한 경우에는 \"이용자\"의 청약철회권이 제한될 수 있습니다.\n  1. 청약의 철회가 불가능한 \"콘텐츠\"에 대한 사실을 표시사항에 포함한 경우\n  2. 시용상품을 제공한 경우\n  3. 한시적 또는 일부이용 등의 방법을 제공한 경우\n  ② \"이용자\"는 다음 각 호의 사유가 있을 때에는 당해 \"콘텐츠\"를 공급받은 날로부터 3월 이내 또는 그 사실을 안 날 또는 알 수 있었던 날부터 30일 이내에 콘텐츠이용계약을 해제·해지할 수 있습니다.\n  1. 이용계약에서 약정한 \"콘텐츠\"가 제공되지 않는 경우\n  2. 제공되는 \"콘텐츠\"가 표시·광고 등과 상이하거나 현저한 차이가 있는 경우\n  3. 기타 \"콘텐츠\"의 결함으로 정상적인 이용이 현저히 불가능한 경우\n  ③ 제1항의 청약철회와 제2항의 계약해제·해지는 \"이용자\"가 전화, 전자우편 또는 모사전송으로 \"회사\"에 그 의사를 표시한 때에 효력이 발생합니다.\n  ④ \"회사\"는 제3항에 따라 \"이용자\"가 표시한 청약철회 또는 계약해제·해지의 의사표시를 수신한 후 지체 없이 이러한 사실을 \"이용자\"에게 회신합니다.\n  ⑤ \"이용자\"는 제2항의 사유로 계약해제·해지의 의사표시를 하기 전에 상당한 기간을 정하여 완전한 \"콘텐츠\" 혹은 서비스이용의 하자에 대한 치유를 요구할 수 있습니다. \n제24조(회사의 계약해제·해지 및 이용제한) ① \"회사\"는 \"회원\"이 제12조 제2항에서 정한 행위를 하였을 경우 사전통지 없이 계약을 해제·해지하거나 또는 기간을 정하여 서비스이용을 제한할 수 있습니다. \n  ② 제1항의 해제·해지는 \"회사\"가 자신이 정한 통지방법에 따라 \"회원\"에게 그 의사를 표시한 때에 효력이 발생합니다.\n  ③ \"회사\"의 해제·해지 및 이용제한에 대하여 \"회원\"은 \"회사\"가 정한 절차에 따라 이의신청을 할 수 있습니다. 이 때 이의가 정당하다고 \"회사\"가 인정하는 경우, \"회사\"는 즉시 서비스의 이용을 재개합니다.\n제25조(회사의 계약해제·해지의 효과) \"회원\"의 귀책사유에 따른 이용계약의 해제·해지의 효과는 제27조를 준용합니다. 다만, \"회사\"는 \"회원\"에 대하여 계약해제·해지의 의사표시를 한 날로부터 7영업일 이내에 대금의 결제와 동일한 방법으로 이를 환급합니다.\n        제5장 피해보상 등\n제26조(콘텐츠하자 등에 의한 이용자피해보상) \"회사\"는 콘텐츠하자 등에 의한 이용자피해보상의 기준·범위·방법 및 절차에 관한 사항을 디지털콘텐츠이용자보호지침에 따라 처리합니다.\n제27조(면책조항) ① \"회사\"는 천재지변 또는 이에 준하는 불가항력으로 인하여 서비스를 제공할 수 없는 경우에는 서비스 제공에 관한 책임이 면제됩니다.  \n  ② \"회사\"는 \"회원\"의 귀책사유로 인한 콘텐츠이용의 장애에 대하여는 책임을 지지 않습니다.  \n  ③ \"회사\"는 \"회원\"이 서비스와 관련하여 게재한 정보, 자료, 사실의 신뢰도, 정확성 등의 내용에 관하여는 책임을 지지 않습니다.  \n  ④ \"회사\"는 \"회원\" 상호간 또는 \"회원\"과 제3자 간에 서비스를 매개로 하여 발생한 분쟁 등에 대하여 책임을 지지 않습니다.\n제28조(분쟁의 해결) \"회사\"는 분쟁이 발생하였을 경우에 \"회원\"이 제기하는 정당한 의견이나 불만을 반영하여 적절하고 신속한 조치를 취합니다. 다만, 신속한 처리가 곤란한 경우에 \"회사\"는 \"회원\"에게 그 사유와 처리일정을 통보합니다.`,
}

var UsageAgreement = map[string]string{
	`개인정보 수집·이용 및 제3자 제공 동의서`: `본인은 다음과 같이 본인의 개인정보를 수집‧이용하고 본 동의서에서 정하는 경우에 한하여 제3자에게 제공하는 것을 동의합니다.
가. (개인정보 수집‧이용자) ㈜아워스포츠네이션
나. (개인정보의 수집‧이용 목적) 서비스 회원가입 및 관리, 재화 또는 서비스 제공, 고충처리
다. 개인정보의 수집‧이용 항목
  ○ 성명, 비밀번호, 휴대폰 번호, 학교, 반, 번호
라. (개인정보의 보유 및 이용기간) 회원 탈퇴시까지. 다만, 관계 법령에 보존 근거가 있는 경우 해당 기간까지 보유. 개인정보처리방침에서 확인 가능.
마. (개인정보 제3자 제공의 경우) 학생의 선생님, 학생의 부모 등 법정대리인, 학생 본인이 허가한 자, 서비스 제공을 위하여 ㈜아워스포츠네이션과 제휴를 맺은 업체에게 제공

  신청인은 개인정보 수집‧이용 및 제3자 제공에 대해 동의를 거부할 수 있으며, 필수항목의 수집에 동의하지 않을 경우 회원가입 신청이 불가합니다.`,
}
