<template>
  <div id="quiz-container">
    <img id="logo-codeify" src="@/assets/codify.svg" alt="headsUP Crown" />
    <h1 id="logo-headline">codeSivity</h1>
    <!-- New Section for User Statistics -->
    <div class="correctAnswers">
      You have
      <strong>{{ correctAnswers }} correct {{ pluralizeAnswer }}!</strong>
    </div>
    <div class="correctAnswers">
      Currently at question {{ index + 1 }} of {{ questions.length }}
    </div>
    <div id="whorelineparent">
      <hr id="whoreline">
    </div>
    <div>
      <h1 v-html="loading ? 'Loading...' : currentQuestion.question"></h1>
      <form v-if="currentQuestion">
        <button
          v-for="answer in currentQuestion.answers"
          :index="currentQuestion.key"
          :key="answer"
          v-html="answer"
          @click.prevent="handleButtonClick"
        ></button>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: "Quiz",
  data() {
    return {
      questions: [],
      loading: true,
      index: 0,
    };
  },
  computed: {
    currentQuestion() {
      if (this.questions !== []) {
        return this.questions[this.index];
      }
      return null;
    },
    score() {
      if (this.questions !== []) {
        // Here, we want to collect data in an object about the users statistics - later be emitted on an event when users finishes quiz
        return {
          allQuestions: this.questions.length,
          answeredQuestions: this.questions.reduce((count, currentQuestion) => {
            if (currentQuestion.userAnswer) {
              // userAnswer is set when user has answered a question, no matter if right or wrong
              count++;
            }
            return count;
          }, 0),
          correctlyAnsweredQuestions: this.questions.reduce(
            (count, currentQuestion) => {
              if (currentQuestion.rightAnswer) {
                // rightAnswer is true, if user answered correctly
                count++;
              }
              return count;
            },
            0
          ),
        };
      } else {
        return {
          allQuestions: 0,
          answeredQuestions: 0,
          correctlyAnsweredQuestions: 0,
        };
      }
    },
    correctAnswers() {
      if (this.questions && this.questions.length > 0) {
        let streakCounter = 0;
        this.questions.forEach(function(question) {
          if (!question.rightAnswer) {
            return;
          } else if (question.rightAnswer === true) {
            streakCounter++;
          }
        });
        return streakCounter;
      } else {
        return "--";
      }
    },
    pluralizeAnswer() {
      // For grammatical correctness
      return this.correctAnswers === 1 ? "Answer" : "Answers";
    },
    quizCompleted() {
      if (this.questions.length === 0) {
        return false;
      }
      /* Check if all questions have been answered */
      let questionsAnswered = 0;
      this.questions.forEach(function(question) {
        question.rightAnswer !== null ? questionsAnswered++ : null;
      });
      return questionsAnswered === this.questions.length;
    },
  },
  watch: {
    quizCompleted(completed) {
      /*
       * Watcher on quizCompleted fires event "quiz-completed"
       * up to parent App.vue component when completed parameter
       * returned by quizCompleted computed property true
       */
      completed &&
        setTimeout(() => {
          this.$emit("quiz-completed", this.score);
        }, 3000); // wait 3 seconds until button animation is over
    },
  },
  methods: {
    async fetchQuestions() {
      this.loading = true;
      let response = await fetch(
        "http://localhost:8090/api/response"
      );
      let jsonResponse = await response.json();
      jsonResponse = jsonResponse.my_response;
      let index = 0; // index is used to identify single answer
      let data = jsonResponse.results.map((question) => {
        // put answers on question into single array
        question.answers = [
          question.correct_answer,
          ...question.incorrect_answers,
        ];
        /* Shuffle question.answers array */
        for (let i = question.answers.length - 1; i > 0; i--) {
          const j = Math.floor(Math.random() * (i + 1));
          [question.answers[i], question.answers[j]] = [
            question.answers[j],
            question.answers[i],
          ];
        }
        // mention in Step 1
        question.rightAnswer = null;
        question.key = index;
        index++;
        return question;
      });
      this.questions = data;
      this.loading = false;
    },
    handleButtonClick: function(event) {
      /* Find index to identiy question object in data */
      let index = event.target.getAttribute("index");
      let pollutedUserAnswer = event.target.innerHTML; // innerHTML is polluted with decoded HTML entities e.g ' from &#039;
      /* Clear from pollution with ' */
      let userAnswer = pollutedUserAnswer.replace(/'/, "&#039;");
      /* Set userAnswer on question object in data */
      this.questions[index].userAnswer = userAnswer;
      /* Set class "clicked" on button with userAnswer -> for CSS Styles; Disable other sibling buttons */
      event.target.classList.add("clicked");
      let allButtons = document.querySelectorAll(`[index="${index}"]`);
      for (let i = 0; i < allButtons.length; i++) {
        if (allButtons[i] === event.target) continue;
        allButtons[i].setAttribute("disabled", "");
      }
      /* Invoke checkAnswer to check Answer */
      this.checkAnswer(event, index);
    },
    checkAnswer: function(event, index) {
      let question = this.questions[index];
      if (question.userAnswer) {
        if (this.index < this.questions.length - 1) {
          setTimeout(
            function() {
              this.index += 1;
            }.bind(this),
            3000
          );
        }
        if (question.userAnswer === question.correct_answer) {
          /* Set class on Button if user answered right, to celebrate right answer with animation joyfulButton */
          event.target.classList.add("rightAnswer");
          /* Set rightAnswer on question to true, computed property can track a streak out of 10 questions */
          this.questions[index].rightAnswer = true;
        } else {
          /* Mark users answer as wrong answer */
          event.target.classList.add("wrongAnswer");
          this.questions[index].rightAnswer = false;
          /* Show right Answer */
          let correctAnswer = this.questions[index].correct_answer;
          let allButtons = document.querySelectorAll(`[index="${index}"]`);
          allButtons.forEach(function(button) {
            if (button.innerHTML === correctAnswer) {
              button.classList.add("showRightAnswer");
            }
          });
        }
      }
    },
  },
  mounted() {
    this.fetchQuestions();
  },
};
</script>

<style scoped>
#whoreline {
  border-top: 4px dotted;
  color: black;
}
#whorelineparent {
  padding-top: 10px;
}
#quiz-container {
  margin: 1rem auto;
  padding: 1rem;
  max-width: 750px;
  
}
#logo-headline {
  font-size: 3rem;
  padding: 0.5rem;
  color: #000000;
  text-align: center;
  letter-spacing: 5px;
}
#logo-codeify {
  display: block;
  width: 40%;
  margin: 0 auto;
}
@media only screen and (max-width: 500px) {
  #logo-crown {
    width: 30%;
  }
  #logo-headline {
    font-size: 1.8rem;
  }
}
h1 {
  font-size: 1.3rem;
  padding: 0.7rem;
  letter-spacing: 5px;
  word-spacing: 17.5px;
  text-align: center;
}
form {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: center;
}
button {
  font-size: 20px;
  font-family: 'Source Code Pro', monospace;
  box-sizing: border-box;
  padding: 1rem;
  margin: 0.3rem;
  width: 47%;
  background-color: rgba(100, 100, 100, 0.3);
  border: none;
  border-radius: 0.4rem;
  box-shadow: 3px 5px 5px rgba(0, 0, 0, 0.2);
}
button:hover:enabled {
  position: relative;
  animation-name: example;
  animation-duration: 2s;
}
button:focus {
  outline: none;
}
button:active:enabled {
  transform: scale(1.05);
}
@keyframes example {
  0%   {background-color:rgb(204, 121, 167); up:0px; top:0px;}
  25%  {background-color:rgb(0, 114, 178); up:0px; top:20px;}
  100% {background-color:rgb(204, 121, 167); up:0px; top:0px;}
}
button.clicked {
  pointer-events: none;
}
button.rightAnswer {
  animation: flashButton;
  animation-duration: 700ms;
  animation-delay: 200ms;
  animation-iteration-count: 3;
  animation-timing-function: ease-in-out;
  color: black;
  background: linear-gradient(
    210deg,
    rgba(0, 158, 115, 0.25),
    rgba(0, 158, 115, 0.5)
  );
}
button.wrongAnswer {
  color: black;
  background: linear-gradient(
    210deg,
    rgba(213, 94, 0, 0.25),
    rgba(213, 94, 0, 0.5)
  );
}
button.showRightAnswer {
  animation: flashButton;
  animation-duration: 700ms;
  animation-delay: 200ms;
  animation-iteration-count: 2;
  animation-timing-function: ease-in-out;
  color: black;
  background: linear-gradient(
    210deg,
    rgba(0, 158, 115, 0.25),
    rgba(0, 158, 115, 0.5)
  );
}
.correctAnswers {
  text-align: center;
  font-size: 20px;
  letter-spacing: 5px;
  word-spacing: 17.5px;
  padding-bottom: 10px;
}
</style>