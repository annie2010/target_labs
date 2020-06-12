<img src="../assets/gophernand.png" align="right" width="128" height="auto"/>

<br/>
<br/>
<br/>

# Prometheus Hangman Lab...

## <img src="../assets/lab.png" width="auto" height="32"/> Your Mission

> We're going to play a hangman game Prometheus Style!
> The game consists of a hangman service and a CLI to submit guesses. The hangman
> service queries a given dictionary to get a list of words for the guess
> word. To play the game, we are going to leverage Prometheus metrics to
> track good/bad guess counts as well as tracking a tally of the win/loose rates.
> Sounds cool?

1. Clone the [Labs Repo](https://github.com/gopherland/labs2)
2. Cd prometheus
3. Instrument the hangman code base and add 2 prometheus metrics to track your
   good and bad guesses (see metrics.go and game.go).
4. Next define a prometheus gauge to track your game results:
   ie +1 for wins and -1 for loss (see metrics.go and tally.go)
5. Install prometheus (see lab template README.md)
6. Configure the scraper to scrape your hangman service on a given port.
7. Start your hangman service
8. Run the provided hangman CLI (cmd/cli/main.go)
9. You can now enjoy the fruits of your labor and try out your guessing skills while watching your game performance in the Prometheus dashboard...

## Commands

1. Download install Prometheus
   1. For OSX use the following command

      ```shell
      cd /tmp
      wget https://github.com/prometheus/prometheus/releases/download/v2.18.0-rc.1/prometheus-2.18.0-rc.1.darwin-amd64.tar.gz
      tar -xvzf /tmp/prometheus-2.18.0-rc.1.darwin-amd64.tar.gz
      # IMPORTANT!! -- Make sure the prometheus binary is on your $PATH
      ```

   2. For other platforms please see [Prometheus Install](https://prometheus.io/download)

1. In one terminal launch prometheus with your custom scraper config file

      ```shell
      prometheus --config.file=config/prom_scraper.yml
      ```

1. Prometheus Dashboard

   ```shell
   open http://localhost:9090/graph
   ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
