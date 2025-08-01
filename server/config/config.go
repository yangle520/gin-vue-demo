package config

type Server struct {
	Zap       Zap      `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis  `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	Mongo     Mongo    `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	System    System   `mapstructure:"system" json:"system" yaml:"system"`
	AutoCode  Autocode `mapstructure:"autocode" json:"autocode" yaml:"autocode"`
	// gorm
	Mysql  Mysql           `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local     Local     `mapstructure:"local" json:"local" yaml:"local"`
	AliyunOSS AliyunOSS `mapstructure:"aliyun-oss" json:"aliyun-oss" yaml:"aliyun-oss"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
