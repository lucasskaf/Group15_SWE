import { Component, OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { MovieGeneratorService } from '../services/movie-generator.service';


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

  constructor(
    private movieGeneratorService: MovieGeneratorService
  ){
    
  }

  ngOnInit(): void {
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

  
}
