import { Component, OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { LoginRegisterService } from '../services/login-register.service';
import { Movie } from '../user-auth/user';


@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  isLoginOpen = false
  isAuthenticated
  username = ""
  message = 'Home Page'
  movieList: Movie[] = []
  userWatchlist: Movie[] = []

  constructor(
    private movieGeneratorService: MovieGeneratorService,
    private loginService: LoginRegisterService
  ){}

  ngOnInit(): void {
    this.getMovie()

    Emmiters.userData.subscribe(
      {
        next: (username : string) => {
          this.username = username
        }
      }
    )

    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.isAuthenticated = auth
          this.loginService.setUserAuth(this.isAuthenticated);
          this.message = 'Hey '
        }
      }
    )


    // this.generatorForm = this.formBuilder.group({
    //   actors: this.actorCtrl,
    //   genre: this.formBuilder.control(''),
    //   rating: this.formBuilder.control(''),
    //   runtime: this.formBuilder.control(''),
    //   provider: this.formBuilder.control('')
    // })
    // this.generatorForm = new FormGroup({
    //   // actors: this.actorCtrl,
    //   genres: new FormControl(''),
    //   minRating: new FormControl(''),
    //   maxRuntime: new FormControl(''),
    //   provider: new FormControl('')
    // })
  }

  getMovie() {
    this.movieGeneratorService.getRandomMovie().subscribe(
      {
        next: (resp) => {
          this.movieList = resp
        },
        error: (err) => {
          console.log('ERROR GENERATING MOVIES')
        }
      }
    )
  }

  setUserAuth(auth: boolean) {
    this.isAuthenticated = auth;
  }
}
