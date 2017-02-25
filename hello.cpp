#include <iostream>
#include <string>
using namespace std;
int main(){
  string s = "Hello world!\n";
  for(auto i: s)
    cout << i;
  
  return 0;
}
