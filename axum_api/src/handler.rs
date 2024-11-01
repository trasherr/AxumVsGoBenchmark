use axum::{
    http::StatusCode, Extension, Json
};
use sea_orm::{prelude::Expr, ActiveModelTrait, ColumnTrait, Condition, DatabaseConnection, EntityTrait, IntoActiveModel, QueryFilter, Set};
use rand::Rng;

use crate::{models::{User, UserModel}, utils::APIError};
pub async fn root(
    Extension(db): Extension<DatabaseConnection>
) -> Result<Json<Vec<UserModel>>,APIError> {

    let users = entity::users::Entity::find().all(&db).await
    .map_err(|err| APIError { message: err.to_string(), status_code:StatusCode::INTERNAL_SERVER_ERROR, error_code: Some(50)})?
    .into_iter()
    .map(|item| UserModel::from(&item))
    .collect();

    Ok(Json(users))
}

pub async fn user_create(
    Extension(db): Extension<DatabaseConnection>,
    Json(payload): Json<crate::models::User>,
) -> Result<Json<UserModel>,APIError> {


   
    let user = entity::users::ActiveModel{
        name: Set(payload.name),
        age: Set(payload.age),
        email: Set(payload.email),
        password: Set(payload.password),
        ..Default::default()
    }.insert(&db)
    .await.map_err(|err| APIError { message: err.to_string(), status_code:StatusCode::INTERNAL_SERVER_ERROR, error_code: Some(50)})?;

    let user = UserModel::from(&user);

    // with a status code of `201 Created`
    Ok(Json(user))
}


pub async fn user_login(
    Extension(db): Extension<DatabaseConnection>,
    Json(payload): Json<crate::models::UserLogin>,
) -> Result<Json<UserModel>,APIError> {

    let user = entity::users::Entity::find().filter(
        Condition::all()
        .add(entity::users::Column::Email.eq(payload.email))
        // .add(entity::users::Column::Password.eq(payload.password))

    ).one(&db).await
    .map_err(|err| APIError { message: err.to_string(), status_code:StatusCode::INTERNAL_SERVER_ERROR, error_code: Some(50)})?
    .ok_or(APIError { message: "User not found".to_string(), status_code:StatusCode::NOT_FOUND, error_code: Some(50)})?;
    
    Ok(Json(UserModel::from(&user)))    
}

pub async fn user_delete(
    Extension(db): Extension<DatabaseConnection>,
    Json(payload): Json<crate::models::UserLogin>,
) -> Result<Json<UserModel>,APIError> {
   
    let user = entity::users::Entity::find().filter(
        Condition::all()
        .add(entity::users::Column::Email.eq(payload.email))
        .add(entity::users::Column::Password.eq(payload.password))

    ).one(&db).await
    .map_err(|err| APIError { message: err.to_string(), status_code:StatusCode::INTERNAL_SERVER_ERROR, error_code: Some(50)})?
    .ok_or(APIError { message: "Not Found".to_owned(), status_code: StatusCode::NOT_FOUND, error_code: Some(44) })?;

    let user_delete_model = UserModel::from(&user); 
    entity::users::Entity::delete_by_id(user.id).exec(&db).await.unwrap();
    
    Ok(Json(user_delete_model))
  
}


pub async fn user_update(
    Extension(db): Extension<DatabaseConnection>,
    Json(payload): Json<crate::models::User>,
) -> Result<Json<UserModel>,APIError> {

    let mut user = entity::users::Entity::find().filter(
        Condition::all()
        .add(entity::users::Column::Email.eq(&payload.email))
        .add(entity::users::Column::Password.eq(&payload.password))

    ).one(&db).await
    .map_err(|err| APIError { message: err.to_string(), status_code:StatusCode::INTERNAL_SERVER_ERROR, error_code: Some(50)})?
    .ok_or(APIError { message: "Not Found".to_owned(), status_code: StatusCode::NOT_FOUND, error_code: Some(44) })?
    .into_active_model();
   
    user.name = Set(payload.name);
    user.age = Set(payload.age);
    user.email = Set(payload.email);
    user.password = Set(payload.password);
       
    let user = user.update(&db).await
    .map_err(|err| APIError { message: err.to_string(), status_code:StatusCode::INTERNAL_SERVER_ERROR, error_code: Some(50)})?;

    let user = UserModel::from(&user);

    Ok(Json(user))
}

pub async fn user_update_age(
    Extension(db): Extension<DatabaseConnection>,
) -> Result<Json<Vec<crate::models::User>>,APIError> {

    let new_age = rand::thread_rng().gen_range(1..100);

    let users: Vec<User> = entity::users::Entity::update_many()
    .col_expr(entity::users::Column::Age, Expr::value(new_age))
    .filter(entity::users::Column::Name.contains("1"))
    .exec_with_returning(&db)
    .await
    .map_err(|err| APIError { message: err.to_string(), status_code:StatusCode::INTERNAL_SERVER_ERROR, error_code: Some(50)})?
    .into_iter()
    .map(|item| User::from(&item))
    .collect();

    Ok(Json(users))
}
