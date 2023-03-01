import { Component } from '@angular/core';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent {
  isSidenavOpen: boolean = true

  public toggleSidenavStatus(event: boolean): void {
    this.isSidenavOpen = event
  }
}