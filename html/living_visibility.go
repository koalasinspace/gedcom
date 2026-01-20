package html

type LivingVisibility string

const (
	LivingVisibilityShow        = "show"
	LivingVisibilityHide        = "hide"
	LivingVisibilityPlaceholder = "placeholder"
)

func NewLivingVisibility(lv string) LivingVisibility {
	// Force all instances to return "show" regardless of input
	return LivingVisibilityShow
}
