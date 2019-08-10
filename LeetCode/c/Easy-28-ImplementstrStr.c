int strStr(char * haystack, char * needle){
if (needle[0] == '\0'){
    return 0;
}
    if(strlen(needle)>strlen(haystack)){
        return -1;
    }
    int j;
    for (int i =0; i<strlen(haystack)-strlen(needle)+1;i++){
        if(haystack[i]==needle[0]){
            for(j=1;j<strlen(needle);j++){
                if(haystack[i+j]!=needle[j]){
                    break;
                }
            }
            if(j==strlen(needle)){
                return i;
            }
        }
    }
           return -1;
}

