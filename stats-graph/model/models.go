package model

var (
	CheckoutsCategory  = &StatCategory{string(StatTypeCheckouts), "Checkouts", &AxisNames{"Time", "Checkouts"}, StatTypeCheckouts, nil}
	DeclinesCategory   = &StatCategory{string(StatTypeDeclines), "Declines", &AxisNames{"Time", "Decline"}, StatTypeDeclines, nil}
	ErrorsCategory     = &StatCategory{string(StatTypeErrors), "Errors", &AxisNames{"Time", "Errors"}, StatTypeErrors, nil}
	FailedCategory     = &StatCategory{string(StatTypeFailed), "Fail", &AxisNames{"Time", "Failures"}, StatTypeFailed, nil}
	TasksCategory      = &StatCategory{string(StatTypeTasksRunning), "Tasks", &AxisNames{"Time", "Tasks Running"}, StatTypeTasksRunning, nil}
	MoneySpentCategory = &StatCategory{string(StatTypeMoneySpent), "Fail", &AxisNames{"Time", "Failures"}, StatTypeMoneySpent, nil}
)

var (
	RecaptchaUsageGensCategory = &StatCategory{"recaptcha_usage", "ReCaptcha Usage", &AxisNames{"Time", "Generations"}, StatTypeRecaptchaUsage, &ReCaptcha}
	ShapeCategory              = &StatCategory{"shape_usage", "Shape Usage", &AxisNames{"Time", "Generations"}, StatTypeCookieGens, &Shape}
	TicketCategory             = &StatCategory{"ticket_usage", "Ticket Usage", &AxisNames{"Time", "Generations"}, StatTypeCookieGens, &Ticket}
	AkamaiGens                 = &StatCategory{"akamai_usage", "Akamai Usage", &AxisNames{"Time", "Generations"}, StatTypeCookieGens, &Akamai}
	PerimeterXGens             = &StatCategory{"perimeterx_usage", "PerimeterX Usage", &AxisNames{"Time", "Generations"}, StatTypeCookieGens, &PerimeterX}
	KasadaGens                 = &StatCategory{"kasada_usage", "Kasada Usage", &AxisNames{"Time", "Generations"}, StatTypeCookieGens, &Kasada}
	ThreatmatrixGens           = &StatCategory{"threatmatrix_usage", "Threatmatrix Usage", &AxisNames{"Time", "Generations"}, StatTypeCookieGens, &Threatmatrix}
)

var (
	Shape        = "Shape"
	Ticket       = "Ticket"
	Akamai       = "Akamai"
	PerimeterX   = "PerimeterX"
	Kasada       = "Kasada"
	Threatmatrix = "Threatmatrix"
	ReCaptcha    = "ReCaptcha"
)

func GetCategories() []*StatCategory {
	return []*StatCategory{
		CheckoutsCategory,
		DeclinesCategory,
		ErrorsCategory,
		FailedCategory,
		TasksCategory,
		MoneySpentCategory,
		RecaptchaUsageGensCategory,
		ShapeCategory,
		TicketCategory,
		AkamaiGens,
		PerimeterXGens,
		KasadaGens,
		ThreatmatrixGens,
	}
}
