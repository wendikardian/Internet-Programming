let name 
let email 
let phone 
let selected
let additional= []
let modal
let form

function submitData(){
    name = document.querySelector("#fullname").value;
    email = document.querySelector("#email").value;
    phone = document.querySelector("#phone").value;
    console.log(name)
    console.log(email)
    console.log(phone)

    let radioInputs = document.querySelectorAll('input[name="package"]');
    
    for(radioInput of radioInputs){
        if(radioInput.checked){
            selected = parseInt(radioInput.value);
            break;
        }
    }
    console.log(selected)

    let checkboxInputs = document.querySelectorAll('input[name="additional"]');
    
    for(checkboxInput of checkboxInputs){
        if(checkboxInput.checked){
            additional.push(checkboxInput.value)
        }
    }
    console.log(additional)
    document.querySelector(".modal").style.display = "block"
}


function display(){
    document.querySelector(".container-form").style.display = "none"
    document.querySelector(".modal").style.display = "none"
    document.querySelector(".result-data").style.display = "block"
    document.querySelector(".name-result").innerHTML = name;
    document.querySelector(".email-result").innerHTML = email;
    document.querySelector(".phone-result").innerHTML = phone;
    totalPayment = 0
    if(selected == 500 ){
        document.querySelector(".selected-result").innerHTML = "Premium Package (England Soccer and more 50 Channels) = Rp. 500.000 ";
        totalPayment += 500
    }else if(selected == 250){
        document.querySelector(".selected-result").innerHTML = "Silver Package (England Soccer and more 25 Channels) = Rp. 250.000 ";
        totalPayment += 250
    }else{
        document.querySelector(".selected-result").innerHTML = "Basic Package (25 Channels) = Rp. 150.000";
        totalPayment +=150
    }
    let ul = document.querySelector(".additional")
    let li
    if(additional.length > 0){
        totalPayment += additional.length*50;
        totalAdded = additional.length*50
        for(item of additional){
            li = document.createElement("li")  
            li.appendChild(document.createTextNode(item))
            ul.appendChild(li)
        }
        document.querySelector('.total-add').innerHTML = "Total Additional = Rp. " +  totalAdded + ".000"
    }else{
        document.querySelector('.total-add').innerHTML = "Rp. 0.-"
    }
    document.querySelector('.total-payment').innerHTML = totalPayment + ".000";
    
}

function again(){
    document.querySelector(".container-form").style.display = "block"
    document.querySelector(".modal").style.display = "none"
    document.querySelector(".result-data").style.display = "none"
    window.location.reload();
}