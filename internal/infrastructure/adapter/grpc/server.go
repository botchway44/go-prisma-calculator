package grpc

import (
	"context"
	"log/slog"

	pb "go-prisma-calculator/generated/proto"
	"go-prisma-calculator/internal/domain/ports/in"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Adapter is the gRPC adapter that connects to our application's core.
type Adapter struct {
	pb.UnimplementedCalculatorServiceServer
	usecase in.CalculatorPort
	logger  *slog.Logger
}

// NewAdapter is the constructor that fx uses to create an instance.
// It receives the application port and logger as dependencies.
func NewAdapter(usecase in.CalculatorPort, logger *slog.Logger) *Adapter {
	return &Adapter{usecase: usecase, logger: logger}
}

// Add handles the gRPC request for the Add RPC.
func (a *Adapter) Add(ctx context.Context, req *pb.AddRequest) (*pb.CalculationResponse, error) {
	a.logger.Info("Handling gRPC Add request", slog.Int("a", int(req.GetA())), slog.Int("b", int(req.GetB())))

	calc, err := a.usecase.Add(ctx, req.GetA(), req.GetB())
	if err != nil {
		a.logger.Error("Usecase failed for gRPC Add", slog.String("error", err.Error()))
		return nil, status.Error(codes.Internal, "an unexpected error occurred")
	}

	a.logger.Info("gRPC Add request successful", slog.Int("result", calc.Result))
	return &pb.CalculationResponse{Result: int32(calc.Result)}, nil
}

// Divide handles the gRPC request for the Divide RPC.
func (a *Adapter) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.CalculationResponse, error) {
	a.logger.Info("Handling gRPC Divide request", slog.Int("dividend", int(req.GetDividend())), slog.Int("divisor", int(req.GetDivisor())))

	calc, err := a.usecase.Divide(ctx, req.GetDividend(), req.GetDivisor())
	if err != nil {
		a.logger.Error("Usecase failed for gRPC Divide", slog.String("error", err.Error()))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	a.logger.Info("gRPC Divide request successful", slog.Int("result", calc.Result))
	return &pb.CalculationResponse{Result: int32(calc.Result)}, nil
}