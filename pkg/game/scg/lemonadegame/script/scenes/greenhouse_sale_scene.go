// Code generated by scg 1, .
//
// LemonadeGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
	base_matchers "github.com/ThCompiler/go_game_constractor/director/matchers"
	"github.com/ThCompiler/go_game_constractor/director/scene"
	"github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/manager"
	"github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/script/errors"
	"github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/script/matchers"
)

const (
	// AgreeGreenhouseSaleButtonText - text for button Agree
	AgreeGreenhouseSaleButtonText = "Ага"
	// BackGreenhouseSaleButtonText - text for button Back
	BackGreenhouseSaleButtonText = "Нет"
)

// GreenhouseSale scene
type GreenhouseSale struct {
	TextManager manager.TextManager
	NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *GreenhouseSale) React(ctx *scene.Context) scene.Command {
	// TODO Write the actions after GreenhouseSale scene has been played
	switch {
	// Buttons select
	case ctx.Request.NameMatched == AgreeGreenhouseSaleButtonText && ctx.Request.WasButton:

	case ctx.Request.NameMatched == BackGreenhouseSaleButtonText && ctx.Request.WasButton:

		// Matcher select
	case ctx.Request.NameMatched == base_matchers.AgreeMatchedString:

	case ctx.Request.NameMatched == matchers.BackMatchedString:

	}

	sc.NextScene = GreenhouseSaleScene // TODO: manually set next scene after reaction
	return scene.NoCommand
}

// Next function returning next scene
func (sc *GreenhouseSale) Next() scene.Scene {
	switch sc.NextScene {
	case ShopScene:
		return &Shop{
			TextManager: sc.TextManager,
		}
	case GreenhouseInfoScene:
		return &GreenhouseInfo{
			TextManager: sc.TextManager,
		}
	}

	return &GreenhouseSale{
		TextManager: sc.TextManager,
	}
}

// GetSceneInfo function returning info about scene
func (sc *GreenhouseSale) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
	var (
		greenhouseCost uint64
	)

	// TODO Write some actions for get data for texts

	text, _ := sc.TextManager.GetGreenhouseSaleText(
		greenhouseCost,
	)
	return scene.Info{
		Text: text,
		ExpectedMessages: []scene.MessageMatcher{
			base_matchers.Agree,
			matchers.BackMatcher,
		},
		Buttons: []scene.Button{
			{
				Title: AgreeGreenhouseSaleButtonText,
			},
			{
				Title: BackGreenhouseSaleButtonText,
			},
		},
		Err: errors.NotUnderstandGreenhouseSaleError,
	}, true
}