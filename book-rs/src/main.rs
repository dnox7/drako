use contracts::author::v1::{GetAuthorRequest, author_service_client::AuthorServiceClient};
use tonic::Request;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = AuthorServiceClient::connect("http://127.0.0.1:8080").await?;
    println!("*** SIMPLE GRPC ***");

    let response = client
        .get_author(Request::new(GetAuthorRequest { id: 1 }))
        .await?;

    println!("RESPONSE (get_author_by_id) = {response:?}");
    Ok(())
}
