package grpc

import (
	"context"

	pb "go-prisma-calculator/generated/proto" // Your generated proto code
	"go-prisma-calculator/internal/domain/ports/in"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Adapter struct {
	pb.UnimplementedCalculatorServiceServer
	usecase in.CalculatorPort
}

func NewAdapter(usecase in.CalculatorPort) *Adapter {
	return &Adapter{usecase: usecase}
}

func (a *Adapter) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	calc, err := a.usecase.Add(ctx, req.GetA(), req.GetB())
	if err != nil {
		return nil, status.Error(codes.Internal, "an unexpected error occurred")
	}
	return &pb.AddResponse{Result: int32(calc.Result)}, nil
}

func (a *Adapter) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.AddResponse, error) {
	calc, err := a.usecase.Divide(ctx, req.GetDividend(), req.GetDivisor())
	if err != nil {
		// Translate domain error to a gRPC error
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.AddResponse{Result: int32(calc.Result)}, nil
}
