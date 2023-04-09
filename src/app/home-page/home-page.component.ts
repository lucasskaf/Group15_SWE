import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { LoginRegisterService } from '../services/login-register.service';
import { FormBuilder } from '@angular/forms';

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

  generatorForm

  constructor(
    private formBuilder: FormBuilder
  ){}

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

    this.generatorForm = this.formBuilder.group({
      actor: this.formBuilder.control(''),
      genre: this.formBuilder.control(''),
      rating: this.formBuilder.control(''),
      runtimehour: this.formBuilder.control(''),
      runtimeminutes: this.formBuilder.control(''),
      provider: this.formBuilder.control('')
    })
  }

  onSubmit(generatorData){
    console.log(generatorData)
  }
}
