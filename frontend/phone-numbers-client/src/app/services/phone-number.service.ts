import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { PhoneNumber } from '../model/phone-number';

const LIMIT = 10

@Injectable({
  providedIn: 'root'
})

export class PhoneNumberService {


  private baseUrl = "http://localhost:8080/api/phone-numbers?limit="+LIMIT;

  constructor(private httpClient: HttpClient) { }

  getPhoneNumbersList(page: number,country: string,state: string): Observable<PhoneNumber[]> {

    const offset = page * LIMIT
    return this.httpClient.get<PhoneNumber[]>(this.baseUrl+"&offset="+offset+"&country="+country+"&state="+state);
  }

}
