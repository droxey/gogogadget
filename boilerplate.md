# Echo API

## Configs

Add a config `struct` for your project:

```go
type Config struct {
	Server struct {
		Host         string `json:"host"`
		Port         string `json:"port"`
		IsProduction bool   `json:"is_production"`
		PasswordSalt []byte `json:"password_salt"`
		JwtSecret    []byte `json:"jwt_secret"`
		Debug        bool   `json:"debug"`
	} `json:"server"`
	Postgres struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DB       string `json:"db"`
	} `json:"postgres"`
}
```

Use a JSON file to configure DB settings:

```json
{
    "server": {
      "host": "0.0.0.0",
      "port": "8080",
      "jwt_secret": "mysecret",
      "password_salt": "mysalted",
      "debug": false
    },
    "postgres": {
      "host": "localhost",
      "port": "5432",
      "user": "backend",
      "password": "mypass",
      "db": "backend"
    }
  }
```

Add the JSON file's specific path in `server.go`:

```go
	configPath := flag.String("config", "./config/production.json", "path of the config file")
	flag.Parse()

	// Read config
	config, err := config.FromFile(*configPath)
	if err != nil {
		log.Fatal(err)
	}
```

Use `configs.GetConfig()` to get configurations, for example: `models/user.go`

```go
func (u *User) SetPassword() {
	cfg := config.GetConfig()
	key  := argon2.Key([]byte(u.Password), cfg.Server.PasswordSalt, 3, 32*1024, 4, 32)
	u.Password = hex.EncodeToString(key)
}
```

## Models

Contains the data structures used for communication between different layers, using gorm.

```go
type User struct {
	gorm.Model
	RoleID       uint   `gorm:"index;not null;default:'2'" json:"role_id,omitempty" valid:"int, required"`
	FirstName    string `gorm:"type:varchar(155);not null" json:"first_name,omitempty" valid:"required"`
	LastName     string `gorm:"type:varchar(155);not null" json:"last_name,omitempty" valid:"required"`
	....
}
```

## Repositories

DRY DB pattern. The base `Repository{}` has common functions like:

```go
func (r *Repository) Find(model interface{}, filter, value string) bool {. .. }

func (r *Repository) Save(model interface{}) bool { ... }
```

To create a new `Repository`:

```go
type UserRepository struct {
	*Repository
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{&Repository{db}}
}

func (ur *UserRepository) FindByCredentials(user *models.User) bool {
	if ur.DB.Where("email = ? and password = ?", user.Email, user.Password).First(&user).RecordNotFound() {
		return false
	}

	return true
}
```

## Resources

API layer that wires up the HTTP routes with the corresponding service APIs.

```go
type UserController struct{}

func (uc *UserController) Register(c echo.Context) error {...}
func (uc *UserController) Login(c echo.Context) error {...}
```

## Routes

Register your in `routes/base.go`

```go
func Init(e *echo.Echo) {
	RegisterUser(e)
}
```

Describe routes for your resources:

```go
func RegisterUser(e *echo.Echo) {
	userController := new(resources.UserController)

	e.POST("auth/register", userController.Register)
	e.POST("auth/login", userController.Login)

}
```

## Middlewares

Custom middlewares belong in `/middlewares`. An example DB middleware to add `db` to `context`:

```go
func DBMiddleware(db *gorm.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			next(c)
			return nil
		}
	}
}
```

## Databases

Manage connection and logic for your databases in `/databases`.

## Helpers

Helper and security functions belong in `/helpers`.
