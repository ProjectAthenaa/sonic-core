package sonic

import (
	"errors"
	module "github.com/ProjectAthenaa/sonic-core/protos"
	"regexp"
)

//All gRPC acceptable statuses redefined as simple variables instead of unreadable capitalized variables
const (
	StatusStarting              = module.STATUS_STARTING
	StatusMonitoring            = module.STATUS_MONITORING
	StatusProductFound          = module.STATUS_PRODUCT_FOUND
	StatusAddingToCart          = module.STATUS_ADDING_TO_CART
	StatusSolvingCaptcha        = module.STATUS_SOLVING_CAPTCHA
	StatusCheckingOut           = module.STATUS_CHECKING_OUT
	StatusCheckedOut            = module.STATUS_CHECKED_OUT
	StatusError                 = module.STATUS_ERROR
	StatusActionNeeded          = module.STATUS_ACTION_NEEDED
	StatusGeneratingCookies     = module.STATUS_GENERATING_COOKIES
	StatusTaskNotFound          = module.STATUS_TASK_NOT_FOUND
	StatusWaitingForCheckout    = module.STATUS_WAITING_FOR_CHECKOUT
	StatusCheckoutError         = module.STATUS_CHECKOUT_ERROR
	StatusCheckoutFailed        = module.STATUS_CHECKOUT_FAILED
	StatusCheckoutDuplicate     = module.STATUS_CHECKOUT_DUPLICATE
	StatusCheckoutOOS           = module.STATUS_CHECKOUT_OOS
	StatusCheckoutDecline       = module.STATUS_CHECKOUT_DECLINE
	StatusCheckoutWaitingFor3DS = module.STATUS_CHECKOUT_WAITING_FOR_3DS
	StatusCheckout3DSError      = module.STATUS_CHECKOUT_3DS_ERROR
)

var (
	redisURLRegex     = regexp.MustCompile(`rediss://\w+:\w+@.*:\d+`)
	channelEmptyError = errors.New("channel_name_cannot_be_empty")
	redisFormatError = errors.New("redis address needs to have correct format")
)
