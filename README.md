 # server-graphql

GOLANG API template's module

## How to use template

You just have to include this line inside a .go file

```
import server_graphql "gitlab.com/fendcer-company/common/golang-modules/server-graphql"
```

You can of course include any submodule of this module.

Then if your compile fails, you should launch this line

```
go env -w GOPRIVATE=gitlab.com/fendcer-company/common/golang-modules/*
```

/!\ Using this needs you the rights to access this repo and you have set up an 
ssh key to your gitlab account

## If you have errors

### Linux

Adding these lines to your ~/.gitonfig:

```
[url "ssh://git@gitlab.com/"]
        insteadOf = https://gitlab.com/

[user]
        email = your.email@format.com
        name = yourName
```

And inside the ~/.netrc:

If you DON'T have the Double Auth actived:

```
machine gitlab.com
        login xxx
        password xxx
```

If you HAVE the Double Auth activated: 

You first need to create an access token Gitlab > Settings > Access Token

Create a token with read only rights (repo/apis/user/registry)

Then insert in ~/.netrc

```
machine gitlab.com
        login xxx
        password you_access_token
```


### Windows

#### Solution 1

You need to change your .gitconfig as well

You can launch this command to localize/see content

```
git config --global -e
```

Then you need to add the following lines
```
[url "git@gitlab.com/"]
        insteadOf = https://gitlab.com/

[user]
        email = your.email@format.com
        name = yourName
```

Try now launching 

```
ssh -T git@gitlab.com
```

If you got 

```
Welcome to Gitlab, @YourUsername
```

It then seems to work, but it not obviously means that you will be able to run the code

#### Solution 2

Still not working ? 

You can try modify ~/.ssh/config like this:

```
# GitLab.com
    Host gitlab.com
    Preferredauthentications publickey
    IdentifyFile ~/.ssh/your_key_name
```

#### Solution 3

If it's still still not working, you might try to add you key to the ssh-agent (if using openSsl)

First make sure ssh-agent is running, for this search in your bottom bar "services"

It opens a GUI Windows, then look for "OpenSSH AUthentication Agent"

If it's unabled, then right click > properties > Type of running: Automatic

Then right click > run

Once you have this ssh-agent running, you can your ssh key

```
ssh-add ~/.ssh/your_ssh_key
```

Launch

```
ssh-add -l
```

You can see your key register, you can try re-running the go projet
