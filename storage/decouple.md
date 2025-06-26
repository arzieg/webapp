1. Create a Domain Model
Create a separate User struct in a package like domain or core, representing the business concept without any GORM annotations:


// domain/user.go
type User struct {
	ID        uint
	FirstName string
	LastName  string
	Emails    []Email
	LastIP    string
}

This becomes the shared abstraction between api and storage.

2. Convert Between Domain and DB Models in the Storage Layer
Keep UserDBModel in your models or storage package. Then write mapper functions to translate between User and UserDBModel:


// storage/mapper.go
func toDBModel(u domain.User) UserDBModel {
	return UserDBModel{
		Model: gorm.Model{ID: u.ID},
		FirstName: u.FirstName,
		LastName: u.LastName,
		// map emails and other fields...
	}
}

func toDomainModel(m UserDBModel) domain.User {
	return domain.User{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		// map emails and other fields...
	}
}

3. Storage Layer Returns Domain Models
That way, your storage code will look like:

func (s *UserStorageImpl) GetUserByID(id uint) (*domain.User, error) {
	var model UserDBModel
	if err := s.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	user := toDomainModel(model)
	return &user, nil
}

Why It’s Worth It
- No GORM leakage into api or domain layers
- Simplifies testing: You can mock behavior without knowing your DB layout
- Future flexibility: If you change your DB or ORM, the rest of your app doesn’t break
