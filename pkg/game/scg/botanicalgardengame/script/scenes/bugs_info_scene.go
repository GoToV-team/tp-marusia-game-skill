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

// BugsInfo scene
type BugsInfo struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *BugsInfo) React(_ *scene.Context) scene.Command {

	sc.NextScene = ShopScene
	return scene.NoCommand
}

// Next function returning next scene
func (sc *BugsInfo) Next() scene.Scene {
	return &Shop{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *BugsInfo) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {

	text, _ := sc.TextManager.GetBugsInfoText()
	return scene.Info{
		Text:             text,
		ExpectedMessages: []scene.MessageMatcher{},
		Buttons:          []scene.Button{},
	}, false
}