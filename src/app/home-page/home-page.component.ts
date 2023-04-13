import { Component, OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { MovieGeneratorService } from '../services/movie-generator.service';
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
  movieList: Movie[] = [
    {
      title: "Movie1",
      release_date: "01-01-2023",
      poster_path: "https://cdn.shopify.com/s/files/1/1057/4964/products/Avengers-Endgame-Vintage-Movie-Poster-Original-1-Sheet-27x41.jpg?v=1670821335",
      user_rating: 10
    },
    {
      title: "Movie2",
      release_date: "01-01-2023",
      poster_path: "https://cdn.shopify.com/s/files/1/1057/4964/products/Avengers-Endgame-Vintage-Movie-Poster-Original-1-Sheet-27x41.jpg?v=1670821335",
      user_rating: 10
    },
    {
      title: "Movie3",
      release_date: "01-01-2023",
      poster_path: "https://cdn.shopify.com/s/files/1/1057/4964/products/Avengers-Endgame-Vintage-Movie-Poster-Original-1-Sheet-27x41.jpg?v=1670821335",
      user_rating: 10
    },
    {
      title: "Movie4",
      release_date: "01-01-2023",
      poster_path: "https://cdn.shopify.com/s/files/1/1057/4964/products/Avengers-Endgame-Vintage-Movie-Poster-Original-1-Sheet-27x41.jpg?v=1670821335",
      user_rating: 10
    },
    {
      title: "Movie5",
      release_date: "01-01-2023",
      poster_path: "https://cdn.shopify.com/s/files/1/1057/4964/products/Avengers-Endgame-Vintage-Movie-Poster-Original-1-Sheet-27x41.jpg?v=1670821335",
      user_rating: 10
    },
    {
      title: "Movie6",
      release_date: "01-01-2023",
      poster_path: "https://cdn.shopify.com/s/files/1/1057/4964/products/Avengers-Endgame-Vintage-Movie-Poster-Original-1-Sheet-27x41.jpg?v=1670821335",
      user_rating: 10
    }
  ]

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
