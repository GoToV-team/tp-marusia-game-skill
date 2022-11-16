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
    base_matchers "github.com/ThCompiler/go_game_constractor/director/scriptdirector/matchers"
    "github.com/ThCompiler/go_game_constractor/director/scriptdirector/scene"
    "github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/manager"
    "github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/script/errors"
    "github.com/evrone/go-clean-template/pkg/game/scg/botanicalgardengame/script/matchers"
)

const (
    // BackDropsSaleButtonText - text for button Back
    BackDropsSaleButtonText = "Нисколько"
)

// DropsSale scene
type DropsSale struct {
    TextManager manager.TextManager
    NextScene   SceneName
}

// React function of actions after scene has been played
func (sc *DropsSale) React(ctx *scene.Context) scene.Command {
    switch {
    // Buttons select
    case ctx.Request.NameMatched == BackDropsSaleButtonText && ctx.Request.WasButton:
        sc.NextScene = ShopScene
        // Matcher select
    case ctx.Request.NameMatched == base_matchers.PositiveNumberMatchedString:
        sc.NextScene = DropsInfoScene
    case ctx.Request.NameMatched == base_matchers.PositiveNumberInWordsMatchedString:
        sc.NextScene = DropsInfoScene
    case ctx.Request.NameMatched == matchers.BackMatchedString:

        sc.NextScene = ShopScene
    }

    return scene.NoCommand
}

// Next function returning next scene
func (sc *DropsSale) Next() scene.Scene {
    switch sc.NextScene {
    case ShopScene:
        return &Shop{
            TextManager: sc.TextManager,
        }
    case DropsInfoScene:
        return &DropsInfo{
            TextManager: sc.TextManager,
        }
    }

    return &DropsSale{
        TextManager: sc.TextManager,
    }
}

// GetSceneInfo function returning info about scene
func (sc *DropsSale) GetSceneInfo(_ *scene.Context) (scene.Info, bool) {

    text, _ := sc.TextManager.GetDropsSaleText()
    return scene.Info{
        Text: text,
        ExpectedMessages: []scene.MessageMatcher{
            base_matchers.PositiveNumberMatcher,
            base_matchers.PositiveNumberInWordsMatcher,
            matchers.BackMatcher,
        },
        Buttons: []director.Button{
            {
                Title: BackDropsSaleButtonText,
            },
        },
        Err: errors.NotUnderstandDropsSaleError,
    }, true
}
