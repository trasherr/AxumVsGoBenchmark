use serde::{Deserialize, Serialize};


#[derive(Serialize,Deserialize)]
pub struct User {
    pub name: String,
    pub age: i32,
    pub email: String,
    pub password: String,
}
impl User {
    pub fn from(user_entity: &entity::users::Model) -> User{
        User { name: user_entity.name.clone(), age: user_entity.age, email: user_entity.email.clone(), password: user_entity.password.clone() }
    }
}

#[derive(Serialize,Deserialize,Debug)]
pub struct UserLogin {
    pub email: String,
    pub password: String,
}

#[derive(Serialize,Deserialize)]
pub struct UserModel {
    pub id: i64,
    pub name: String,
    pub age: i32,
    pub email: String,
    pub password: String,
}

impl UserModel {
    pub fn from(user_entity: &entity::users::Model) -> UserModel{
        UserModel { id: user_entity.id, name: user_entity.name.clone(), age: user_entity.age, email: user_entity.email.clone(), password: user_entity.password.clone() }
    }
}