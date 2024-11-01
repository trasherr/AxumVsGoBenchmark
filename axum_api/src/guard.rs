
use axum::{extract::Request, http::StatusCode, middleware::Next, response::Response};

use crate::utils::APIError;

pub async fn guard(req: Request, next: Next) -> Result<Response, APIError> {

    let token = req.headers().get("Authorization")
    .ok_or(APIError { message: "Unauthorized".to_owned(), status_code: StatusCode::UNAUTHORIZED, error_code: Some(41)  })?
    .to_str()
    .map_err(|err| APIError { message: err.to_string(), status_code: StatusCode::UNAUTHORIZED, error_code: Some(41)  })?
    .to_owned();
    let _token = token.replace("Bearer", "");
    if token.is_empty()  {
        return Err(APIError { message: "Unauth".to_string(), status_code:StatusCode::UNAUTHORIZED, error_code: Some(41)});
    }

    Ok(next.run(req).await)
}