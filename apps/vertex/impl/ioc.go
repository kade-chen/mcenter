package impl

import (
	"context"
	"iter"

	"github.com/rs/zerolog"
	"google.golang.org/genai"

	"github.com/kade-chen/library/exception"
	"github.com/kade-chen/library/ioc"
	"github.com/kade-chen/library/ioc/config/log"
	"github.com/kade-chen/mcenter/apps/vertex"
	"github.com/kade-chen/mcenter/apps/vertex/provider"
)

var _ vertex.ServiceGemini3 = (*service)(nil)

func init() {
	ioc.Controller().Registry(&service{})
}

type service struct {
	// col *mongo.Collection
	// token.UnimplementedRPCServer
	ioc.ObjectImpl
	log *zerolog.Logger

	// policy  policy.Service
	// ns      namespace.Service
	// checker security.Checker
	// domain  domain.Service
	// notify  notify.Service
}

func (s *service) Init() error {
	// dc := ioc_mongo.DB().Collection(s.Name())
	// indexs := []mongo.IndexModel{
	// 	{
	// 		Keys:    bson.D{{Key: "refresh_token", Value: -1}},
	// 		Options: options.Index().SetUnique(true),
	// 	},
	// 	{
	// 		Keys: bson.D{{Key: "issue_at", Value: -1}},
	// 	},
	// }

	// _, err := dc.Indexes().CreateMany(context.Background(), indexs)
	// if err != nil {
	// 	return err
	// }

	// s.col = dc

	s.log = log.Sub(s.Name())
	// s.ns = ioc.Controller().Get(namespace.AppName).(namespace.Service)
	// s.policy = ioc.Controller().Get(policy.AppName).(policy.Service)
	// s.domain = ioc.Controller().Get(domain.AppName).(domain.Service)
	// s.notify = ioc.Controller().Get(notify.AppName).(notify.Service)

	// s.checker, err = security.NewChecker()
	// if err != nil {
	// 	return fmt.Errorf("new checker error, %s", err)
	// }

	// 初始化所有的auth provider
	if err := provider.Init(); err != nil {
		return err
	}
	return nil
}

func (service) Name() string {
	return vertex.AppName
}

func (service) NoStreamingGenerateContent(ctx context.Context, gemini2setting *vertex.Gemini2Config) (*genai.GenerateContentResponse, error) {
	if gemini2setting == nil {
		gemini2setting = vertex.NewDefaultsGemini2Config()
	}
	// gemini2setting.ModelName = "gemini-1.0-pro"
	a, exists := vertex.GRANT_MODEL_value[gemini2setting.ModelName]
	if !exists {
		return nil, exception.NewBadRequest("Model Name: %s , And cannot be empty or no in the registry", gemini2setting.ModelName)
	}

	pr, _ := provider.ProviderRegistry()
	if l := provider.CheckProviderRegistry(a, pr); !l {
		return nil, exception.NewNotFound("Model Name: %s does not exist", gemini2setting.ModelName)
	}

	// impl, err := provider.GetModelIssuer(vertex.GRANT_MODEL_GEMINI_2_0_Flash)
	// impl.NoStreamingGenerateContent(context.Background(), gemini2setting)

	return provider.GetModelIssuer(a).NoStreamingGenerateContent(context.Background(), gemini2setting)
}

func (service) StreamingGenerateContent(ctx context.Context, gemini2setting *vertex.Gemini2Config) (iter.Seq2[*genai.GenerateContentResponse, error], error) {

	if gemini2setting == nil {
		gemini2setting = vertex.NewDefaultsGemini2Config()
	}
	// gemini2setting.ModelName = "gemini-1.0-pro"
	a, exists := vertex.GRANT_MODEL_value[gemini2setting.ModelName]
	if !exists {
		return nil, exception.NewBadRequest("Model Name: %s , And cannot be empty or no in the registry", gemini2setting.ModelName)
	}

	pr, _ := provider.ProviderRegistry()
	if l := provider.CheckProviderRegistry(a, pr); !l {
		return nil, exception.NewNotFound("Model Name: %s does not exist", gemini2setting.ModelName)
	}

	return provider.GetModelIssuer(a).StreamingGenerateContent(context.Background(), gemini2setting), nil
}
