// Code generated by scg 1, .
//
// BotanicalGardenGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/manager"
)

// Rules scene
type Rules struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *Rules) React(_ *scene.Context) scene.Command {
	return scene.ApplyStashedScene
}

// Next function returning next scene
func (sc *Rules) Next() scene.Scene {

	return &Rules{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *Rules) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {

	text, _ := sc.TextManager.GetRulesText()
	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []scene.Button{},
	}, false
}