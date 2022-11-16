// Code generated by scg 1, .
//
// BotanicalGardenGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
    "github.com/ThCompiler/go_game_constractor/director"
    "github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
    "github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/manager"
)

// NewDay scene
type NewDay struct {
    TextManager manager.TextManager
    NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *NewDay) React(_ *scene.Context) scene.Command {
    return scene.NoCommand
}

// Next function returning next scene
func (sc *NewDay) Next() scene.Scene {
    return &Shop{
        TextManager: sc.TextManager,
    }
}

// GetSceneInfo function returning info about scene
func (sc *NewDay) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {

    text, _ := sc.TextManager.GetNewDayText()
    return scene.Info{
        Text:             text,
        ExpectedMessages: []scene.MessageMatcher{},
        Buttons:          []director.Button{},
    }, false
}
