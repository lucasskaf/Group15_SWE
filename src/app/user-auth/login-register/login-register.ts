export function loginInteraction() {
    const signUpButton: HTMLElement | null = document.getElementById('signUp');
    const signInButton: HTMLElement | null = document.getElementById('signIn');
    const container: HTMLElement | null = document.getElementById('container');

    if(signUpButton != null && signInButton != null && container != null){
      signUpButton.addEventListener('click', () => {
        container.classList.add("right-panel-active");
      });
      
      signInButton.addEventListener('click', () => {
        container.classList.remove("right-panel-active");
      });
    }
}