// Code generated by scg 1, .
//
// LemonadeGame-SceneStructs
//
// Command:
// scg
//.

package scenes

import (
    "github.com/ThCompiler/go_game_constractor/director"
    base_matchers "github.com/ThCompiler/go_game_constractor/director/scriptdirector/matchers"
    "github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
    "github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/manager"
    "github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/script/errors"
    "github.com/evrone/go-clean-template/pkg/game/scg/lemonadegame/script/matchers"
)

const (
    // AgreeFlySaleButtonText - text for button Agree
    AgreeFlySaleButtonText = "Ага"
    // BackFlySaleButtonText - text for button Back
    BackFlySaleButtonText = "Нет"
)

// FlySale scene
type FlySale struct {
    TextManager manager.TextManager
    NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *FlySale) React(ctx *scene.Context) scene.Command {
    // TODO Write the actions after FlySale scene has been played
    switch {
    // Buttons select
    case ctx.Request.NameMatched == AgreeFlySaleButtonText && ctx.Request.WasButton:

    case ctx.Request.NameMatched == BackFlySaleButtonText && ctx.Request.WasButton:

        // Matcher select
    case ctx.Request.NameMatched == base_matchers.AgreeMatchedString:

    case ctx.Request.NameMatched == matchers.BackMatchedString:

    }

    sc.NextScene = FlySaleScene // TODO: manually set next scene after reaction
    return scene.NoCommand
}

// Next function returning next scene
func (sc *FlySale) Next() scene.Scene {
    switch sc.NextScene {
    case ShopScene:
        return &Shop{
            TextManager: sc.TextManager,
        }
    case FlyInfoScene:
        return &FlyInfo{
            TextManager: sc.TextManager,
        }
    }

    return &FlySale{
        TextManager: sc.TextManager,
    }
}

// GetSceneInfo function returning info about scene
func (sc *FlySale) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
    var (
        flyCost uint64
    )

    // TODO Write some actions for get data for texts

    text, _ := sc.TextManager.GetFlySaleText(
        flyCost,
    )
    return scene.Info{
        Text: text,
        ExpectedMessages: []scene.MessageMatcher{
            base_matchers.Agree,
            matchers.BackMatcher,
        },
        Buttons: []director.Button{
            {
                Title: AgreeFlySaleButtonText,
            },
            {
                Title: BackFlySaleButtonText,
            },
        },
        Err: errors.NotUnderstandFlySaleError,
    }, true
}
