use axum::{
    routing::{delete, get, post, put}, Extension, Router, middleware
};
use sea_orm::{Database, DatabaseConnection};

mod models;
mod handler;
mod guard;
mod utils;
mod db_initializer;

#[tokio::main]
async fn main() {

    // build our application with a route
    let db: DatabaseConnection = Database::connect("sqlite://benchmark.db?mode=rwc").await.unwrap();
    // db_initializer::seed(&db).await;

    let app = Router::new()
        
        // `GET /` goes to `root`
        .route("/users", put(handler::user_update_age))
        .route("/user", put(handler::user_update))
        .route("/user", delete(handler::user_delete))
        .route_layer(middleware::from_fn(guard::guard))
        
        .route("/user", get(handler::user_login))
        .route("/user", post(handler::user_create))
        .route("/", get(handler::root))
        .layer(Extension(db));

    // run our app with hyper, listening globally on port 3000
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}


