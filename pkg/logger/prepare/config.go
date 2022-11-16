package prepare

import "github.com/ThCompiler/go_game_constractor/pkg/logger"

type Config struct {
    Level                    logger.LogLevel `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
    LogDir                   string          `env-required:"true" yaml:"log_dir,omitempty,omitempty" env:"LOG_DIR"`
    UseStdAndFIle            bool            `env-required:"true" yaml:"use_std_and_file,omitempty" env:"USER_STD_AND_FILE"`
    AddLowPriorityLevelToCmd bool            `env-required:"true" yaml:"add_low_priority_level_to_cmd,omitempty" env:"ADD_LPL_TO_CMD"`
}