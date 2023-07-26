// Import necessary modules
use reqwest::Result;
use tokio::runtime::Builder;
use tokio::task;

// Async function to send HTTP GET request
async fn send_request(url: String) -> Result<()> {

    // Create a reqwest client
    let client = reqwest::Client::new();

    // Perform the request
    let response = client.get(&url).send().await?;

    // Ensure the response was successful
    if response.status().is_success() {
        // Use the response here if needed, for example:
        let body = response.text().await?;
        println!("Request ok: {}", &url);
    } else {
        println!("Request failed with status: {}", response.status());
    }

    Ok(())
}

async fn say_world() {
    println!("world");
}

#[tokio::main]
async fn main() {
    let handle = tokio::spawn(async {
        // Do some async work
        "return value"
    });


    let out = handle.await.unwrap();
    println!("GOT {}", out);

    let o0 = tokio::spawn(send_request("https://aws.com".to_owned()));
    let o1 = tokio::spawn(send_request("https://jd.com".to_owned()));
    let o2 = tokio::spawn(send_request("https://qq.com".to_owned()));
    let o3 = tokio::spawn(send_request("https://163.com".to_owned()));

    let out = o0.await.unwrap();
    println!("GOT {:?}", out);
}