use rand::Rng;
use sea_orm::{DatabaseConnection, Set,ActiveModelTrait};

pub async fn seed(db: &DatabaseConnection) {

    for i in 0..10000 {
        let new_age = rand::thread_rng().gen_range(1..100);
        
        entity::users::ActiveModel{
            name: Set(format!("test{}",i)),
            age: Set(new_age),
            email: Set(format!("test{}@gmail.com",i)),
            password: Set(format!("password{}",i)),
            ..Default::default()
        }.insert(db)
        .await.unwrap();
    }


}