var myNav = document.querySelector('header');
window.onscroll = function () { 
    console.log(document.body.scrollTop)
    if (window.scrollY > 300) {
        myNav.classList.add("white-navbar");
        myNav.classList.remove("color-navbar");
    } else {
        myNav.classList.add("color-navbar");
        myNav.classList.remove("white-navbar");
    }
};

var slides = document.querySelectorAll('.slide');
var buttons = document.querySelectorAll('.slider-btn')
let currentSlide = 1

var manualNav = function(manual){
    slides.forEach(function(slide) {
        slide.classList.remove('active');
        buttons.forEach(btn => {
            btn.classList.remove('active')
        });
    });
    slides[manual].classList.add('active');
    buttons[manual].classList.add('active');
}

buttons.forEach(function(btn, i) {
    btn.addEventListener('click', function(){
        manualNav(i);
        currentSlide = i;
    })
})


var repeat = function(activeClass){
    let active = document.getElementsByClassName('active');
    let i=1;
    var repeater = function(){
        setTimeout(function() {
            [...active].forEach(function(activeSlide) {
                activeSlide.classList.remove('active')
            });

            slides[i].classList.add('active')
            buttons[i].classList.add('active')
            i++;
            if(slides.length == i){
                i = 0
            }
            if(i>= slides.length){
                return
            }
            repeater();
        }, 5000)
    }
    repeater()
}

repeat()

AOS.init();