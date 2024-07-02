div p-0 
  {{ template "navbar_exit" . }}
  div relative w-full 
    div rounded-tl-lg rounded-tr-lg bg-gradient-to-b from-a-top to-a-bottom h-1/2 w-full
    div rounded-bl-lg rounded-br-lg bg-yellow-300 h-1/2 w-full
    canvas id=board absolute top-0 left-0 w-full h-full 
