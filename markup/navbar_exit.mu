{{ define "navbar_exit" }}
  div navbar bg-base-200 font-familjen
    div navbar-start 
      div btn btn-ghost text-4xl
        a href=/
          img rounded-lg src=logo.png w-12
    div navbar-center flex hidden md:block
    div navbar-end
      div hidden md:block flex space-x-3
        a href=/ link link-hover
          Exit
  {{ end }}
