# Database

## Testing

Current testing database:

- holds ~1000 randomly generated users in the users/ table
- holds ~200 randomly generated doctors in doctors/, with a few assigned countries + specialties

## Caveats

Username is unique, email is unique for every entry

- if a user tries to register as an existing email, error out
- if a user tries to register an existing username, error out

ID only set for users, which is used internally in the session mananger only

- Doctors cannot update from their profile page. They can do so only via correspondence
