//go:generate go run github.com/dmarkham/enumer -type Job -trimprefix Job -transform upper -sql -json -output job_string.go
package macro

import "fmt"

type Job int

const (
	_ Job = iota
	JobGLD
	JobPGL
	JobMRD
	JobLNC
	JobARC
	JobCNJ
	JobTHM
	JobCRP
	JobBSM
	JobARM
	JobGSM
	JobLTW
	JobWVR
	JobALC
	JobCUL
	JobMIN
	JobBTN
	JobFSH
	JobPLD
	JobMNK
	JobWAR
	JobDRG
	JobBRD
	JobWHM
	JobBLM
	JobACN
	JobSMN
	JobSCH
	JobROG
	JobNIN
	JobMCH
	JobDRK
	JobAST
	JobSAM
	JobRDM
	JobBLU
	JobGNB
	JobDNC
	JobRPR
	JobSGE
)

var _jobNameMap = map[Job]string{
	JobGLD: "剑术",
	JobPGL: "格斗",
	JobMRD: "斧术",
	JobLNC: "枪术",
	JobARC: "弓箭",
	JobCNJ: "幻术",
	JobTHM: "咒术",
	JobCRP: "刻木",
	JobBSM: "锻铁",
	JobARM: "铸甲",
	JobGSM: "雕金",
	JobLTW: "制革",
	JobWVR: "裁衣",
	JobALC: "炼金",
	JobCUL: "烹调",
	JobMIN: "采矿",
	JobBTN: "园艺",
	JobFSH: "捕鱼",
	JobPLD: "骑士",
	JobMNK: "武僧",
	JobWAR: "战士",
	JobDRG: "龙骑",
	JobBRD: "诗人",
	JobWHM: "白魔",
	JobBLM: "黑魔",
	JobACN: "秘术",
	JobSMN: "召唤",
	JobSCH: "学者",
	JobROG: "双剑",
	JobNIN: "忍者",
	JobMCH: "机工",
	JobDRK: "黑骑",
	JobAST: "占星",
	JobSAM: "武士",
	JobRDM: "赤魔",
	JobBLU: "青魔",
	JobGNB: "绝枪",
	JobDNC: "舞者",
	JobRPR: "镰刀",
	JobSGE: "贤者",
}

var _jobFullNameMap = map[Job]string{
	JobGLD: "剑术师",
	JobPGL: "格斗家",
	JobMRD: "斧术师",
	JobLNC: "枪术师",
	JobARC: "弓箭手",
	JobCNJ: "幻术师",
	JobTHM: "咒术师",
	JobCRP: "刻木匠",
	JobBSM: "锻铁匠",
	JobARM: "铸甲匠",
	JobGSM: "雕金匠",
	JobLTW: "制革匠",
	JobWVR: "裁衣匠",
	JobALC: "炼金术士",
	JobCUL: "烹调师",
	JobMIN: "采矿工",
	JobBTN: "园艺工",
	JobFSH: "捕鱼人",
	JobPLD: "骑士",
	JobMNK: "武僧",
	JobWAR: "战士",
	JobDRG: "龙骑士",
	JobBRD: "吟游诗人",
	JobWHM: "白魔法师",
	JobBLM: "黑魔法师",
	JobACN: "秘术师",
	JobSMN: "召唤师",
	JobSCH: "学者",
	JobROG: "双剑师",
	JobNIN: "忍者",
	JobMCH: "机工士",
	JobDRK: "暗黑骑士",
	JobAST: "占星术士",
	JobSAM: "武士",
	JobRDM: "赤魔法师",
	JobBLU: "青魔法师",
	JobGNB: "绝枪战士",
	JobDNC: "舞者",
	JobRPR: "钐镰客",
	JobSGE: "贤者",
}

func (i Job) Name() string {
	return _jobNameMap[i]
}

func (i Job) FullName() string {
	return _jobFullNameMap[i]
}

func JobNameString(name string) (Job, error) {
	for i, n := range _jobNameMap {
		if n == name {
			return i, nil
		}
	}
	return Job(0), fmt.Errorf("%s does not belong to Role values", name)
}

func JobFullNameString(fullName string) (Job, error) {
	for i, n := range _jobFullNameMap {
		if n == fullName {
			return i, nil
		}
	}
	return Job(0), fmt.Errorf("%s does not belong to Role values", fullName)
}

func (i Job) Values() []string {
	return JobStrings()
}
