import { Component, EventEmitter, Output, Input } from '@angular/core';

@Component({
  selector: 'app-add-popup-movie',
  templateUrl: './add-popup-movie.component.html',
  styleUrls: ['./add-popup-movie.component.css']
})
export class AddMoviePopupComponent {
  @Output() isClosed = new EventEmitter();
  @Input() isOpen : boolean = false;

  ngOnInit() {
    
  }

  togglePopup() {
    this.isOpen = !this.isOpen;
    this.isClosed.emit(this.isOpen);
  }
}
