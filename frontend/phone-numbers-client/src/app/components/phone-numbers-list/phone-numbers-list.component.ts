import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { PhoneNumber } from 'src/app/model/phone-number';
import { PhoneNumberService } from 'src/app/services/phone-number.service';

@Component({
  selector: 'app-phone-numbers-list',
  templateUrl: './phone-numbers-list.component.html',
  styleUrls: ['./phone-numbers-list.component.css']
})
export class PhoneNumbersListComponent implements OnInit,OnDestroy {

  private country: string = ""
  private state: string = ""
  private page: number = 0
  private phoneNumbersSubscription: Subscription = new Subscription;
  phoneNumbers: PhoneNumber[] = [];
  displayedColumns: string[] = ['country', 'state', 'code', 'number'];

  constructor(private phoneNumberService: PhoneNumberService) { }
 
  ngOnInit(): void {
    this.listPhoneNumbers();
  }

  ngOnDestroy(): void {
    this.phoneNumbersSubscription.unsubscribe()
  }

  listPhoneNumbers() {
    this.phoneNumbersSubscription = this.phoneNumberService.getPhoneNumbersList(this.page,this.country,this.state).subscribe(
      data => {
        this.phoneNumbers = data;
        console.log(this.phoneNumbers)
      }
    )
  }

  onPrev(){
    if (this.page>0){
      this.page--;
      this.phoneNumbersSubscription.unsubscribe()
      this.listPhoneNumbers();
    }
  }
  onNext(){
    if (this.phoneNumbers.length > 0) {
      this.page++;
      this.phoneNumbersSubscription.unsubscribe()
      this.listPhoneNumbers();
    }
  }
  onSelectCountry(data: any){
    this.page=0;
    this.country = data.target.value
    this.phoneNumbersSubscription.unsubscribe()
    this.listPhoneNumbers();
  }

  onSelectState(data: any){
    this.page=0;
    this.state = data.target.value
    this.phoneNumbersSubscription.unsubscribe()
    this.listPhoneNumbers();
  }


    
}
