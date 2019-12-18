const encryption_button = document.getElementById('encryption');
const decryption_button = document.getElementById('decryption');
const encrypted_text = document.getElementById('encrypted-text');
const decrypted_text = document.getElementById('decrypted-text');
const url = window.location.href;

encryption_button.onclick = () => {
    const encrypt_text = document.getElementById('encrypt-text').value;
    fetch(`${url}api/encryption?text=${encrypt_text}`)
        .then(res => {
            return res.json();
        })
        .then(json => {
            encrypted_text.style.color = 'black';
            encrypted_text.innerHTML = json.encrypted_text;
        })
        .catch(err => {
            console.error(err);
        });
};

decryption_button.onclick = () => {
    const decrypt_text = document.getElementById('decrypt-text').value;
    fetch(`${url}api/decryption?text=${decrypt_text}`)
        .then(res => {
            return res.json();
        })
        .then(json => {
            decrypted_text.style.color = 'black';
            decrypted_text.innerHTML = json.decrypted_text;
        })
        .catch(err => {
            console.error(err)
        });
};
