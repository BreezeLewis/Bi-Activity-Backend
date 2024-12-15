// response/errors/student_error/error_helper.go
package student_error

import "errors"

func GetErrorMsg(code int) string {
    if msg, ok := errMsgMap[code]; ok {
        return msg
    }
    return "未知错误"
}

// GetErrorCode 获取错误码
func GetErrorCode(err error) int {
    switch {
    // 基础错误
    case errors.Is(err, ErrStudentNotFoundError):
        return ErrStudentNotFound
    case errors.Is(err, ErrInvalidStudentIDError):
        return ErrInvalidStudentID
    
    // 学院相关错误
    case errors.Is(err, ErrCollegeNotFoundError):
        return ErrCollegeNotFound
    case errors.Is(err, ErrStudentNoCollegeError):
        return ErrStudentNoCollege
    case errors.Is(err, ErrCollegeListNotFoundError):
        return ErrCollegeListNotFound
    
    // 安全设置相关错误
    case errors.Is(err, ErrPasswordIncorrectError):
        return ErrPasswordIncorrect
    case errors.Is(err, ErrPhoneExistsError):
        return ErrPhoneExists
    case errors.Is(err, ErrEmailExistsError):
        return ErrEmailExists
    case errors.Is(err, ErrAccountNotFoundError):
        return ErrAccountNotFound
    case errors.Is(err, ErrInvalidCodeError):
        return ErrInvalidCode
    case errors.Is(err, ErrThirdPartyBoundError):
        return ErrThirdPartyBound
    case errors.Is(err, ErrPhoneRequiredError):
        return ErrPhoneRequired
    case errors.Is(err, ErrPasswordNotMatchError):
        return ErrPasswordNotMatch

    // 活动相关错误
    case errors.Is(err, ErrActivityNotFoundError):
        return ErrActivityNotFound
    case errors.Is(err, ErrInvalidActivityIDError):
        return ErrInvalidActivityID
    case errors.Is(err, ErrActivityStatusInvalidError):
        return ErrActivityStatusInvalid
    case errors.Is(err, ErrActivityFullError):
        return ErrActivityFull
    case errors.Is(err, ErrActivityExpiredError):
        return ErrActivityExpired
    case errors.Is(err, ErrActivityNotStartedError):
        return ErrActivityNotStarted
    case errors.Is(err, ErrActivityFinishedError):
        return ErrActivityFinished
    case errors.Is(err, ErrActivityAuditingError):
        return ErrActivityAuditing
    case errors.Is(err, ErrActivityRejectedError):
        return ErrActivityRejected

    // 活动相关错误
    case errors.Is(err, ErrActivityNotFoundError):
        return ErrActivityNotFound
    case errors.Is(err, ErrInvalidActivityIDError):
        return ErrInvalidActivityID
    case errors.Is(err, ErrActivityStatusInvalidError):
        return ErrActivityStatusInvalid
    case errors.Is(err, ErrActivityFullError):
        return ErrActivityFull
    case errors.Is(err, ErrActivityExpiredError):
        return ErrActivityExpired
    case errors.Is(err, ErrActivityNotStartedError):
        return ErrActivityNotStarted
    case errors.Is(err, ErrActivityFinishedError):
        return ErrActivityFinished
    case errors.Is(err, ErrActivityAuditingError):
        return ErrActivityAuditing
    case errors.Is(err, ErrActivityRejectedError):
        return ErrActivityRejected

    // 参与者相关错误
    case errors.Is(err, ErrParticipantNotFoundError):
        return ErrParticipantNotFound
    case errors.Is(err, ErrParticipantInvalidError):
        return ErrParticipantInvalid

    default:
        return -1
    }
}