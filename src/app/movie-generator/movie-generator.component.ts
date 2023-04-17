import { Component } from '@angular/core';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { FormBuilder, Validators } from '@angular/forms';
import { Movie, User, moviePosts } from '../user-auth/user';
import { OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';

@Component({
  selector: 'bb-movie-generator',
  templateUrl: './movie-generator.component.html',
  styleUrls: ['./movie-generator.component.css']
})
export class MovieGeneratorComponent implements OnInit {

  username
  generatorForm
  postForm
  generatedMovie: Movie = {vote_average: 0.0}
  isPopupOpen: boolean = false
  isAuthenticated: boolean = false
  userWatchedList: Movie[] = []
  isMovieWatched: boolean = false
  isPostOpen: boolean = false

  constructor(private movieGeneratorService : MovieGeneratorService,
    private formBuilder: FormBuilder) {
    this.generatorForm = this.formBuilder.group({
      actors: this.formBuilder.control(''),
      genres: this.formBuilder.control(''),
      min_rating: this.formBuilder.control(''),
      max_runtime: this.formBuilder.control(''),
      streaming_providers: this.formBuilder.control('')
    })
  }

  ngOnInit(): void {
    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.isAuthenticated = auth
        }
      }
    )

    Emmiters.watchList.subscribe(
      {
        next: (userList: Movie[]) => {
          this.userWatchedList = userList
        }
      }
    )

    Emmiters.userData.subscribe(
      {
        next: (username: string) => {
          this.username = username
        }
      }
    )
  }

  getGeneratedMovie(filters){
    this.movieGeneratorService.getRandomMovieWithFilters(filters).subscribe(
      {
        next: (respMovie) => {
          console.log(respMovie)
          this.generatedMovie = respMovie

          if(this.userWatchedList.includes(this.generatedMovie)){
            this.isMovieWatched = true
          }
          else {
            this.isMovieWatched = false
          }

          console.log(`WAS MOVIE WATCHED: ${this.isMovieWatched}`)
          this.showMovie()

          console.log(`GETTING USERNAME: ${this.username}`)

          this.postForm = this.formBuilder.group({
            movieid: this.generatedMovie.id,
            username: this.username,
            title: this.formBuilder.control(''),
            body: this.formBuilder.control('', Validators.required)
          })
        },
        error: (err) => {
          console.log(err)
        }
      }
    )
  }

  showMovie() {
    this.isPopupOpen = true
  }

  closeMovie() {
    console.log('CLICKED ON CLOSE')
    this.isPopupOpen = false
  }

  onSubmit(generatorData){
    console.log(`BFORE: ${generatorData}`)

    // Text process
    if(generatorData.actors){
      generatorData.actors = generatorData.actors.toLowerCase()
    }
    
    let actors_array = generatorData.actors.split(', ')
    
    generatorData.actors = actors_array

    //Rating & Runtime process
    if(generatorData.min_rating) {
      generatorData.min_rating = (parseFloat(generatorData.min_rating) * 2.0)
    }

    if(generatorData.max_runtime) {
      generatorData.max_runtime = parseInt(generatorData.max_runtime)
    }

    console.log(`AFTER: ${generatorData}`)

    this.getGeneratedMovie(generatorData)
  }

  onSubmitPost(post: moviePosts){
    post.movieid = post.movieid.toString()
    console.log(post)

    this.movieGeneratorService.addMoviePosts(post).subscribe(
      {
        next: (resp) => {
          console.log('ADDED SUCCESSFULLY')
        },
        error: (err) => {
          console.log('NOT ADD POST')
        }
      }
    )
  }

  formatLabel(value: number): string {
    if(value % 60 == 0){
      return (value / 60) + 'hrs'
    }

    return (Math.floor(value / 60)) + ':' + (value % 60)
  }

  addToWatchlist(movie: Movie) {
    this.movieGeneratorService.addToWatchList(movie).subscribe(
      {
        next: (resp) => {
          console.log(resp)
          this.isMovieWatched = true
        },
        error: (err) => {
          console.log(err)
        }
      }
    )
  }

  toogleAddPost(){
    this.isPostOpen = !this.isPostOpen
  }
}
