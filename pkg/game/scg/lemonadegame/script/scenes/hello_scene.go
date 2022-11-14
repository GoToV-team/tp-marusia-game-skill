// Code generated by scg 1, .
//
// LemonadeGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/manager"
)

// Hello scene
type Hello struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *Hello) React(_ *scene.Context) scene.Command {
	return scene.NoCommand
}

// Next function returning next scene
func (sc *Hello) Next() scene.Scene {
	return &StartRules{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *Hello) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {

	// TODO Write some actions for get data for texts

	text, _ := sc.TextManager.GetHelloText()
	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []scene.Button{},
	}, false
}