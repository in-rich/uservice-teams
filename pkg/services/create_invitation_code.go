package services

import (
	"context"
	"crypto/rand"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/in-rich/uservice-teams/pkg/dao"
	"github.com/in-rich/uservice-teams/pkg/entities"
	"github.com/in-rich/uservice-teams/pkg/models"
	"math/big"
	"time"
)

const InvitationCodeLength = 6

const InvitationCodeTryCount = 3

var InvitationCodeChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var InvitationCodeCharsLength = big.NewInt(int64(len(InvitationCodeChars)))

type CreateInvitationCodeService interface {
	Exec(ctx context.Context, in *models.CreateInvitationCodeRequest) (*models.CreateInvitationCodeResponse, error)
}

type createInvitationCodeServiceImpl struct {
	createInvitationCodeRepository dao.CreateInvitationCodeRepository
}

func (s *createInvitationCodeServiceImpl) generateUniqueCode() (string, error) {
	output := make([]byte, InvitationCodeLength)

	for i := 0; i < InvitationCodeLength; i++ {
		index, err := rand.Int(rand.Reader, InvitationCodeCharsLength)
		if err != nil {
			return "", errors.Join(ErrGenerateCode, err)
		}

		output[i] = InvitationCodeChars[index.Int64()]
	}

	return string(output), nil
}

func (s *createInvitationCodeServiceImpl) Exec(ctx context.Context, in *models.CreateInvitationCodeRequest) (*models.CreateInvitationCodeResponse, error) {
	var err error
	var res *entities.InvitationCode

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(in); err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	expiresAt := time.Now().Add(168 * time.Hour)

	teamID, err := uuid.Parse(in.TeamID)
	if err != nil {
		return nil, errors.Join(ErrInvalidData, err)
	}

	// Generate candidate codes. If e get an error because of a duplicate, we try with another one.
	for i := 0; i < InvitationCodeTryCount; i++ {
		code, err := s.generateUniqueCode()
		if err != nil {
			return nil, err
		}

		res, err = s.createInvitationCodeRepository.CreateInvitationCode(ctx, uuid.New(), &dao.CreateInvitationCodeData{
			Code:      code,
			TeamID:    teamID,
			ExpiresAt: expiresAt,
		})

		if errors.Is(err, dao.ErrCodeAlreadyExists) {
			continue
		}

		if err != nil {
			return nil, errors.Join(ErrGenerateCode, err)
		}

		return &models.CreateInvitationCodeResponse{
			Code: res.Code,
		}, nil
	}

	return nil, errors.Join(ErrGenerateCode, err)
}

func NewCreateInvitationCodeService(createInvitationCodeRepository dao.CreateInvitationCodeRepository) CreateInvitationCodeService {
	return &createInvitationCodeServiceImpl{
		createInvitationCodeRepository: createInvitationCodeRepository,
	}
}
