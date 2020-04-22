package opslevel

// OpsLevel 开放代理套餐种类
type OpsLevel = string

const (
	// NORMAL 普通套餐
	NORMAL OpsLevel = "normal"
	// VIP VIP套餐
	VIP OpsLevel = "vip"
	// SVIP SVIP套餐
	SVIP OpsLevel = "svip"
	// ENT 专业版套餐
	ENT OpsLevel = "ent"
)
