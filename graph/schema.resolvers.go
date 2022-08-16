package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"learnGraphql/graph/generated"
	"learnGraphql/graph/model"
)

// remember this
//type Resolver struct {
//
//}

// CreateBio is the resolver for the CreateBio field.
func (r *mutationResolver) CreateBio(ctx context.Context, input model.AddBio) (*model.BioData, error) {
	bio, err := r.BioRepository.CreatePerson(&input)
	createdBio := &model.BioData{
		Name: bio.Name,
		Age:  bio.Age,
	}
	if err != nil {
		return nil, err
	}
	return createdBio, nil
}

// DeleteBio is the resolver for the DeleteBio field.
func (r *mutationResolver) DeleteBio(ctx context.Context, age int) (string, error) {
	err := r.BioRepository.DeletePerson(age)
	if err != nil {
		return "", err
	}
	successMessage := "successfully deleted the bio"
	return successMessage, nil
}

// UpdateBio is the resolver for the UpdateBio field.
func (r *mutationResolver) UpdateBio(ctx context.Context, age int) (string, error) {
	//err := r.BioRepository.UpdatePerson(&input, age)
	//if err != nil {
	//	return "nil", err
	//}
	//successMessage := "successfully updated"
	//
	//return successMessage, nil
	panic(fmt.Errorf("not implemented"))

}

// GetBio is the resolver for the GetBio field.
func (r *queryResolver) GetBio(ctx context.Context, id int) (*model.BioData, error) {
	bio, err := r.BioRepository.GetOnePerson(id)
	selectedBio := &model.BioData{
		Name: bio.Name,
		Age:  bio.Age,
	}
	if err != nil {
		return nil, err
	}
	return selectedBio, nil
}

// GetBios is the resolver for the GetBios field.
func (r *queryResolver) GetBios(ctx context.Context) ([]*model.BioData, error) {
	//panic(fmt.Errorf("not implemented"))
	bios, err := r.BioRepository.GetAllPersons()
	if err != nil {
		return nil, err
	}
	return bios, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
