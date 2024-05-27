package models

type RobotMapping struct {
	Model
	TRobotConfID            int    `json:"t_robot_conf_id"`
	TokenID                 int    `json:"token_id"`
	Host                    string `json:"host"`
	Port                    int    `json:"port"`
	FeishuAppID             string `json:"feishu_app_id"`
	FeishuAppSecret         string `json:"feishu_app_secret"`
	FeishuVerificationToken string `json:"feishu_verification_token"`
	FeishuEncryptKey        string `json:"feishu_encrypt_key"`
	FeishuBotName           string `json:"feishu_bot_name"`
	DingtalkClientID        string `json:"dingtalk_client_id"`
	DingtalkClientSecret    string `json:"dingtalk_client_secret"`
	Mode                    int    `json:"mode"`
	DingtalkAesKey          string `json:"dingtalk_aes_key"`
	DingtalkToken           string `json:"dingtalk_token"`
	DingtalkWebhook         string `json:"dingtalk_webhook"`
}

func (RobotMapping) TableName() string {
	return "t_robot_mapping"
}

func GetRobotMappingByFeishuToken(token string) (*RobotMapping, error) {
	var result RobotMapping
	if err := db.Where("feishu_verification_token = ? and deleted_on = 0", token).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func GetRobotMappingByDingtalkWebhook(dingtalkWebhook string) (*RobotMapping, error) {
	var result RobotMapping
	if err := db.Where("dingtalk_webhook = ? and deleted_on = 0", dingtalkWebhook).First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
