const encryption_button = document.getElementById('encryption');
const decryption_button = document.getElementById('decryption');

encryption_button.onclick = () => {
    const encrypt_text = document.getElementById('encrypt-text').value;
    fetch(`http://localhost:8080/api/encryption?text=${encrypt_text}`)
        .then(res => {
            return res.json();
        })
        .then(json => {
            document.getElementById('encrypted-text').innerHTML = json.encrypted_text;
            console.log(json.encrypted_text);
        })
        .catch(err => {
            console.error(err);
        });
};

decryption_button.onclick = () => {
    const decrypt_text = document.getElementById('decrypt-text').value;
    fetch(`http://localhost:8080/api/decryption?text=${decrypt_text}`)
        .then(res => {
            return res.json();
        })
        .then(json => {
            document.getElementById('decrypted-text').innerHTML = json.decrypted_text;
            console.log(json.decrypted_text);
        })
        .catch(err => {
            console.error(err)
        });
};
