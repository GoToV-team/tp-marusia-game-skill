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
    // AgreeApproveEndDayButtonText - text for button Agree
    AgreeApproveEndDayButtonText = "Да"
    // BackApproveEndDayButtonText - text for button Back
    BackApproveEndDayButtonText = "Нет"
)

// ApproveEndDay scene
type ApproveEndDay struct {
    TextManager manager.TextManager
    NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *ApproveEndDay) React(ctx *scene.Context) scene.Command {
    // TODO Write the actions after ApproveEndDay scene has been played
    switch {
    // Buttons select
    case ctx.Request.NameMatched == AgreeApproveEndDayButtonText && ctx.Request.WasButton:

    case ctx.Request.NameMatched == BackApproveEndDayButtonText && ctx.Request.WasButton:

        // Matcher select
    case ctx.Request.NameMatched == base_matchers.AgreeMatchedString:

    case ctx.Request.NameMatched == matchers.BackMatchedString:

    }

    sc.NextScene = ApproveEndDayScene // TODO: manually set next scene after reaction
    return scene.NoCommand
}

// Next function returning next scene
func (sc *ApproveEndDay) Next() scene.Scene {
    switch sc.NextScene {
    case EndOfDayScene:
        return &EndOfDay{
            TextManager: sc.TextManager,
        }
    case ShopScene:
        return &Shop{
            TextManager: sc.TextManager,
        }
    }

    return &ApproveEndDay{
        TextManager: sc.TextManager,
    }
}

// GetSceneInfo function returning info about scene
func (sc *ApproveEndDay) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {

    // TODO Write some actions for get data for texts

    text, _ := sc.TextManager.GetApproveEndDayText()
    return scene.Info{
        Text: text,
        ExpectedMessages: []scene.MessageMatcher{
            base_matchers.Agree,
            matchers.BackMatcher,
        },
        Buttons: []director.Button{
            {
                Title: AgreeApproveEndDayButtonText,
            },
            {
                Title: BackApproveEndDayButtonText,
            },
        },
        Err: errors.NotUnderstandApproveEndDayError,
    }, true
}
