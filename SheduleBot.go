package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand/v2"
	"strings"
)

const TOKEN = "7647905012:AAGhEF9pQvplTQtmHRoVuumOfKz7te3UEZw"

var bot *tgbotapi.BotAPI
var chatId int64

var scheduleBotNames = [3]string{"стивен", "стив", "кинг"}

var answers = []string{
	"Иногда в жизни важно отпустить прошлое, чтобы сделать шаг вперёд",
	"Беги от того, что разрушает тебя, даже если всем вокруг кажется, что это невозможно",
	"Страх — худший советчик, не давай ему решать за себя",
	"Порой нужно сделать то, чего ты боишься, ведь только так ты узнаешь, кто ты есть",
	"Иногда лучше не заглядывать в тёмные тайники души — но если решишься, будь готов встретить себя настоящего",
	"Месть редко приносит облегчение. Лучше попытаться простить, чем тащить груз обиды",
	"Не слушай тех, кто говорит, что ты недостаточно хорош — докажи это делом",
	"Бывают моменты, когда правда важнее дружбы",
	"Не все двери нужно открывать, даже если ты безумно любопытен",
	"Иногда беспокойство за будущее превращается в худший кошмар — лучше действовать сейчас",
	"Бывает, что чудовища лишь в твоей голове, но справиться с ними никто, кроме тебя, не сможет",
	"Смерть — не всегда худшее, что может случиться",
	"С настоящей любовью нельзя бороться — но порой стоит отпустить",
	"Иногда только крайняя боль толкает нас меняться",
	"Дорога возникает под ногами идущего — остановка куда страшнее провала",
	"Ожидание может убить веру быстрее, чем самый страшный удар",
	"Мир не делится на чёрное и белое — принимай решения сердцем",
	"Иногда проще поверить в невозможное, чем в очевидное",
	"В трудную минуту важно не отдаляться, а искать поддержку, даже если кажется, что её нет",
	"Твой главный враг — это ты сам",
	"Бывает, что прошлое призывает нас назад — но не всегда стоит туда возвращаться",
	"Не жди идеального момента, он не наступит — двигайся сейчас",
	"Не бойся потерять, бойся не попробовать",
	"Иногда цена за желание слишком высока — подумай, прежде чем чего-то страстно захотеть",
	"Секреты рано или поздно всплывают — лучше быть честным сразу",
	"Иногда нужно просто замолчать и дать жизни идти своим чередом",
	"Не пытайся контролировать всё — случай может сыграть за тебя",
	"Труднее всего прощать самого себя — но иначе не освободишься",
	"Быть другим — не значит быть хуже",
	"Джентльменство в мелочах говорит о большом человеке",
}

func connectWithTelegram() {
	var err error
	if bot, err = tgbotapi.NewBotAPI(TOKEN); err != nil {
		panic("Cannot connect to Telegram")
	}
}

func sendMessage(msg string) {
	msgConfig := tgbotapi.NewMessage(chatId, msg)
	bot.Send(msgConfig)
}

func isMessageForFortuneTeller(update *tgbotapi.Update) bool {
	if update.Message == nil || update.Message.Text == "" {
		return false
	}

	msgInLowerCase := strings.ToLower(update.Message.Text)
	for _, name := range scheduleBotNames {
		if strings.Contains(msgInLowerCase, name) {
			return true
		}
	}
	return false
}

func getScheduleBotAnswer() string {
	index := rand.IntN(len(answers))
	return answers[index]
}

func sendAnswer(update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatId, getScheduleBotAnswer())
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}

func main() {
	connectWithTelegram()

	updateConfig := tgbotapi.NewUpdate(0)
	for update := range bot.GetUpdatesChan(updateConfig) {
		if update.Message != nil && update.Message.Text == "/start" {
			chatId = update.Message.Chat.ID
			sendMessage("Задай мне свой вопрос, назвав меня по имени." +
				"Ответом на вопрос должен быть \"Да\" либо \"Нет\". Например, \"Кинг, я готов сменить работу?\"" +
				"либо \"Стивен, я действительно хочу отправиться на эту вечеринку?\"")
		}
		if isMessageForFortuneTeller(&update) {
			sendAnswer(&update)
		}
	}
}
