int lengthOfLastWord(char * s){
    int i = strlen(s)-1;
    int out = 0;
    bool charFound = false;
    while(i>=0){
        if(s[i]==' '){
            if (charFound) {
                break;
            }else{
                i--;
            }
        }else{
            i--;
            out++;
            charFound = true;
        }
    }
    return out;  
        
}
