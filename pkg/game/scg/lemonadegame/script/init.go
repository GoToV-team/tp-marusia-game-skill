// Code generated by scg 1,  DO NOT EDIT .
//
// LemonadeGame-Director config
//
// Command:
// scg
// DO NOT EDIT .

package script

import (
    game "github.com/ThCompiler/go_game_constractor/director/scriptdirector"
    "github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/manager"
    "github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/script/scenes"
)

const GoodByeCommand = "Пока"

func NewLemonadeGameScript(manager manager.TextManager) game.SceneDirectorConfig {
    return game.SceneDirectorConfig{
        StartScene: &scenes.Hello{
            TextManager: manager,
        },
        GoodbyeScene: &scenes.Goodbye{
            TextManager: manager,
        },
        EndCommand: GoodByeCommand,
    }
}
