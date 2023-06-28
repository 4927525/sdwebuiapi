package test

import (
	"fmt"
	"regexp"
	"testing"
	"unicode"
)

func TestCharset(t *testing.T) {
	s1 := "(extremely detailed cg unity 8k wallpaper,masterpiece,best quality,ultra-detailed),(best illumination,best shadow,an extremely delicate and beautiful),classic,(impasto,photorealistic,),dynamic angle,floating,finely detail,(bloom),(shine),glinting stars,beautiful detailed girl,extremely delicate and beautiful girls,beautiful detailed eyes,glowing eyes,blank stare,beautiful face,extremely_beautiful_detailed_anime_face,cute face,bright skin,{{{{{solo}}}}},((starry sky)),star river,array stars,holy,noble,intricate light,dynamic hair,haircut,dynamic fuzziness,beautiful and aesthetic,intricate light,manga and anime,girl=1,1girl:2.0<{bare shoulders,diamond and glaring eyes,medium chest,beautiful detailed cold face,black hair/black pupil,wavy hair,white silk stocking,(((hanfu))),headphones around neck,bare navel,{{skinny legs}},floating hair,{{{{{wet}}}}},{{{{wet clothes}}}},{{{see-through raincoat}}},heart in eye,heart-shaped pupils,{underwear},{take a shower},((reflective eyes)),((hair dripping)),water eyes,drunk,light blush,looking_at_viewer,intricate detail,{{solo}},{translucent open navel dress made of tulle},barefoot sandals,{ribbon made of tulle},{delicate skin},beautiful detailed eyes,{{shed tears}},pick dyeing,{wet clothes}},on water<{(the lake has a girl's reflection:1.2),(reflection:0.85),cherry_blossoms,evening,falling_leaves,forest,ginkgo_leaf,gradient_sky,holding_leaf,leaf,long_sleeves,maple_leaf,molten_rock,nature,orange_flower,outdoors,petals,river,scenery,sky,standing,sunset,tree,twilight,water,wisteria,(cinematiclighting),dramatic angle,depth of field,(upper body on water:1.6)},underwater<{((sparkle background)),(((lower body under water:1.5))),coral,(((tyndall effect))),(underwater forest),(sunlight),floating hair,(beautiful detailed water,bubbles,beautiful and detailed bubbles,beautiful and detailed oceans,beautiful and detailed corals,corals,seaweeds,sea beds,gravels,{ top-down light },{ light tracing },{dim light},beautiful and detailed water,ray refraction,dream like benthos,transparent fish,{purple glowing jellyfish},pearl,gemstone,{trapped in bubbles},ocean bottom,tropical fish,kentaurosu,fairey swordfish,clownfish,seaweed,{dreamy},magic array,magic jewel,)}"
	fmt.Println(IsChineseChar(s1))
}

// 或者封装函数调用
func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}
