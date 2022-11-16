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
    base_matchers "github.com/ThCompiler/go_game_constractor/director/scriptdirector/matchers"
    "github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/manager"
    "github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/script/errors"
    "github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/script/matchers"
)

const (
    // AgreeBugsSaleButtonText - text for button Agree
    AgreeBugsSaleButtonText = "Ага"
    // BackBugsSaleButtonText - text for button Back
    BackBugsSaleButtonText = "Нет"
)

// BugsSale scene
type BugsSale struct {
    TextManager manager.TextManager
    NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *BugsSale) React(ctx *scene.Context) scene.Command {
    switch {
    // Buttons select
    case ctx.Request.NameMatched == AgreeBugsSaleButtonText && ctx.Request.WasButton:
        sc.NextScene = BugsInfoScene
    case ctx.Request.NameMatched == BackBugsSaleButtonText && ctx.Request.WasButton:
        sc.NextScene = ShopScene
        // Matcher select
    case ctx.Request.NameMatched == base_matchers.AgreeMatchedString:
        sc.NextScene = BugsInfoScene
    case ctx.Request.NameMatched == matchers.BackMatchedString:
        sc.NextScene = ShopScene
    }

    return scene.NoCommand
}

// Next function returning next scene
func (sc *BugsSale) Next() scene.Scene {
    switch sc.NextScene {
    case ShopScene:
        return &Shop{
            TextManager: sc.TextManager,
        }
    case BugsInfoScene:
        return &BugsInfo{
            TextManager: sc.TextManager,
        }
    }

    return &BugsSale{
        TextManager: sc.TextManager,
    }
}

// GetSceneInfo function returning info about scene
func (sc *BugsSale) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {
    var (
        bugsCost uint64
    )

    text, _ := sc.TextManager.GetBugsSaleText(
        bugsCost,
    )
    return scene.Info{
        Text: text,
        ExpectedMessages: []scene.MessageMatcher{
            base_matchers.Agree,
            matchers.BackMatcher,
        },
        Buttons: []director.Button{
            {
                Title: AgreeBugsSaleButtonText,
            },
            {
                Title: BackBugsSaleButtonText,
            },
        },
        Err: errors.NotUnderstandBugsSaleError,
    }, true
}
